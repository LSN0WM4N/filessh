package sshclient

import (
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

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
