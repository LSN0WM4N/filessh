package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

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
	homeDir := os.Getenv("HOME")

	if homeDir == "" || runtime.GOOS == "windows" {
		fmt.Printf(
			"It seems like you are not using the best OS in the world.\n" +
				"I am so sorry but i still working in linux main version, be patient. There\n" +
				"will be a windows version soon. (or not).",
		)
		os.Exit(1)
	}

	exp := explorer.New(homeDir)
	p := tea.NewProgram(cli.NewModel(exp))

	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
