package main

import (
	"flag"
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	cli "github.com/LSN0WM4N/filessh/pkg/cli"
	explorer "github.com/LSN0WM4N/filessh/pkg/explorer"
)

func main() {
	// Load flag values
	forceBackward := flag.Bool(
		"force-backward",
		false,
		"Force backward-compatible terminal output",
	)
	flag.Parse()
	cli.SetForceBackward(*forceBackward)

	// Create the explorer
	exp := explorer.New("/home/sn0wm4n")
	p := tea.NewProgram(cli.NewModel(exp))

	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
