package cli

import (
	"github.com/LSN0WM4N/filessh/pkg/explorer"
)

type model struct {
	explorer *explorer.Explorer
	entries  []explorer.Entry
	cursor   int

	width  uint8
	height uint8
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

func (m model) computeVisibleWindow(listHeight int) (entries []explorer.Entry, start int, selected *explorer.Entry) {
	total := len(m.entries)
	end := total

	if total > listHeight {
		end = listHeight

		if m.cursor > CURSOR_OFFSET {
			start = m.cursor - CURSOR_OFFSET
			end = start + listHeight
		}

		if end > total {
			end = total
			start = end - listHeight
		}
	}

	if m.cursor >= 0 && m.cursor < total {
		selected = &m.entries[m.cursor]
	}

	return m.entries[start:end], start, selected
}
