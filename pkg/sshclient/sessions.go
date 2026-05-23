package sshclient

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/LSN0WM4N/filessh/pkg/bus"
	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

/*
PTY Mode - Shell mode, in this mode you will able to work
exactly as a normal shell, no TUI, no render, just a shell
*/
func PTYMode(session *ssh.Session, ctx context.Context, bus *bus.EventBus) (*ssh.Session, error) {
	fd := int(os.Stdin.Fd())

	width, height, err := term.GetSize(fd)
	if err != nil {
		width, height = 80, 24
	}

	oldState, err := term.MakeRaw(fd)
	if err != nil {
		fmt.Printf("Failed to set raw terminal: %v\n", err)
		return nil, err
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
		return nil, err
	}

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	if err := session.Shell(); err != nil {
		fmt.Printf("Failed to start shell: %v\n", err)
		return nil, err
	}

	go watchWindowResize(session, fd)

	if err := session.Wait(); err != nil {
		if err.Error() != "EOF" {
			fmt.Printf("\nSession ended: %v\n", err)
		}
	}

	return session, nil
}

/*
Pipe Mode - TUI Mode, in this mode the view content is all
rendered by the TUI engine, in this mode I plan to give you
access to some basic commands, not a whole shell

!IMPORTANT: In this mode, programs such as `vim`, `htop` or
these like will no work, for these use PTY mode
*/
func PipeMode(session *ssh.Session, ctx context.Context, bus *bus.EventBus) (*ssh.Session, error) {
	fd := int(os.Stdin.Fd())

	width, height, err := term.GetSize(fd)
	if err != nil {
		width, height = 80, 24
	}

	oldState, err := term.MakeRaw(fd)
	if err != nil {
		fmt.Printf("Failed to set raw terminal: %v\n", err)
		return nil, err
	}
	defer term.Restore(fd, oldState)

	_ = session
	_ = bus

	drawConsoleFrame(width, height)
	fmt.Printf("\x1b[2;3Hstill working on it")
	fmt.Printf("\x1b[%d;1H", height)

	buf := make([]byte, 1)
	for {
		n, err := os.Stdin.Read(buf)
		if err != nil || n == 0 {
			break
		}
		if buf[0] == 'q' || buf[0] == 'Q' || buf[0] == 3 {
			break
		}
	}

	fmt.Print("\x1b[0m\x1b[2J\x1b[H")

	return session, nil
}

func drawConsoleFrame(width, height int) {
	if width < 2 || height < 2 {
		return
	}

	horizontal := "+" + strings.Repeat("-", width-2) + "+"
	blank := "|" + strings.Repeat(" ", width-2) + "|"

	fmt.Print("\x1b[2J")
	for row := 1; row <= height; row++ {
		fmt.Printf("\x1b[%d;1H", row)
		if row == 1 || row == height {
			fmt.Print(horizontal)
		} else {
			fmt.Print(blank)
		}
	}
}
