package tui

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func TermSize() WindowSize {
	fd := int(os.Stdout.Fd())

	width, height, err := term.GetSize(fd)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Cols: %d\n", width)
	fmt.Printf("Rows: %d\n", height)

	return WindowSize{
		Width:  width,
		Height: height,
	}
}
