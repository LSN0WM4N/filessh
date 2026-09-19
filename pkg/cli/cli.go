package cli

import (
	"github.com/LSN0WM4N/filessh/pkg/explorer"
)

type model struct {
	explorer *explorer.Explorer
	entries  []explorer.Entry
	cursor   int
}

type entriesLoadedMsg struct {
	entries []explorer.Entry
}

func NewModel(exp *explorer.Explorer) model {
	return model{
		explorer: exp,
		entries:  make([]explorer.Entry, 0),
		cursor:   0,
	}
}
