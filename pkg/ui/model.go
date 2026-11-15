package ui

import (
	"slices"

	tea "charm.land/bubbletea/v2"
	"github.com/LSN0WM4N/filessh/pkg/events"
	"github.com/LSN0WM4N/filessh/pkg/explorer"
	"github.com/LSN0WM4N/filessh/pkg/ui/components/filelist.go"
	"github.com/LSN0WM4N/filessh/pkg/ui/components/preview.go"
	"github.com/LSN0WM4N/filessh/pkg/ui/theme"
)

type Model struct {
	explorer *explorer.Explorer
	config   *theme.Config

	filelist filelist.Model
	preview  preview.Model

	width  uint8
	height uint8
}

func New(exp *explorer.Explorer) Model {
	config := theme.New(theme.Options{})
	entries, _ := exp.List(false, true)

	var first *explorer.Entry
	if len(entries) > 0 {
		first = &entries[0]
	}

	return Model{
		explorer: exp,
		config:   &config,

		filelist: filelist.New(&config, &entries),
		preview:  preview.New(&config, first),
	}
}

func (m Model) Init() tea.Cmd {
	return func() tea.Msg {
		list, err := m.explorer.List(true, true)
		slices.SortFunc(list, compareByTypeFirst)

		if err != nil {
			return "Unable to load the current dir"
		}

		return events.SetupExplorer{
			Entries: &list,
		}
	}
}
