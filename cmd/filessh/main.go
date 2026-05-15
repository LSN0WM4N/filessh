package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/LSN0WM4N/filessh/pkg/tui"
)

func main() {
	size := tui.TermSize()
	fmt.Printf("Cols: %d\nRows: %d\n", size.Width, size.Height)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGWINCH)

	for range sig {
		size := tui.TermSize()
		fmt.Printf("Cols: %d\nRows: %d\n", size.Width, size.Height)
	}

	fmt.Println("[+] Still working on it, sorry :(")
}
