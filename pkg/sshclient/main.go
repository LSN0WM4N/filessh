package sshclient

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/crypto/ssh"
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
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	err := session.RequestPty("xterm", 40, 120, modes)
	if err != nil {
		panic(err)
	}

	// stdin, err := session.StdinPipe()
	// if err != nil {
	// 	panic(err)
	// }

	// stdout, err := session.StdoutPipe()
	// if err != nil {
	// 	panic(err)
	// }

	// stderr, err := session.StderrPipe()
	// if err != nil {
	// 	panic(err)
	// }

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	err = session.Shell()
	if err != nil {
		panic(err)
	}
	session.Shell()

	return session
}

func ReadOutput(stdout io.Reader) {
	io.Copy(os.Stdout, stdout)
}
