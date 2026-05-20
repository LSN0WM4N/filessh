package sshclient

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

func SetupConnection(user UserConfig) (*ssh.Client, error) {
	passwordPtr := user.Password
	var auth []ssh.AuthMethod

	if passwordPtr != nil {
		auth = append(auth, ssh.Password(*passwordPtr))
	}

	config := &ssh.ClientConfig{
		User:            *user.Username,
		Auth:            auth,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return ssh.Dial("tcp", fmt.Sprintf("%s:%s", user.Host, user.Port), config)
}

func OpenSession(client *ssh.Client) (*ssh.Session, error) {
	return client.NewSession()
}

func SetupShell(session *ssh.Session) *ssh.Session {
	fd := int(os.Stdin.Fd())

	width, height, err := term.GetSize(fd)
	if err != nil {
		width, height = 80, 24
	}

	oldState, err := term.MakeRaw(fd)
	if err != nil {
		fmt.Printf("Failed to set raw terminal: %v\n", err)
		return session
	}
	defer term.Restore(fd, oldState)

	termType := os.Getenv("TERM")
	if termType == "" {
		termType = "xterm-256color"
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	if err := session.RequestPty(termType, height, width, modes); err != nil {
		fmt.Printf("Failed to request PTY: %v\n", err)
		return session
	}

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	if err := session.Shell(); err != nil {
		fmt.Printf("Failed to start shell: %v\n", err)
		return session
	}

	go watchWindowResize(session, fd)

	if err := session.Wait(); err != nil {
		if err.Error() != "EOF" {
			fmt.Printf("\nSession ended: %v\n", err)
		}
	}

	return session
}

func watchWindowResize(session *ssh.Session, fd int) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGWINCH)
	defer signal.Stop(sigCh)

	for range sigCh {
		w, h, err := term.GetSize(fd)
		if err == nil {
			_ = session.WindowChange(h, w)
		}
	}
}
