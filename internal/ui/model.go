package ui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/LSN0WM4N/filessh/internal/events"
	"github.com/LSN0WM4N/filessh/internal/explorer"
	"github.com/LSN0WM4N/filessh/internal/ui/components/filelist"
	"github.com/LSN0WM4N/filessh/internal/ui/components/preview"
	"github.com/LSN0WM4N/filessh/internal/ui/theme"
)

type Model struct {
	explorer *explorer.Explorer
	config   *theme.Config

	filelist filelist.Model
	preview  preview.Model

	width  int
	height int
}

func New(exp *explorer.Explorer) Model {
	config := theme.New(theme.Options{})
	entries, _ := exp.List(true, true)

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

		if err != nil {
			return "Unable to load the current dir"
		}

		return events.SetupExplorer{
			Entries: &list,
		}
	}
}
