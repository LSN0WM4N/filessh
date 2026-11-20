package filelist

import (
	tea "charm.land/bubbletea/v2"

	"github.com/LSN0WM4N/filessh/internal/events"
	"github.com/LSN0WM4N/filessh/internal/explorer"
	"github.com/LSN0WM4N/filessh/internal/ui/theme"
)

type Model struct {
	config  *theme.Config
	entries *[]explorer.Entry

	width  int
	height int

	startCursor int
	endCursor   int
	cursor      int
}

const VIEW_GAP float32 = 0.55

func New(
	config *theme.Config,
	entries *[]explorer.Entry,
) Model {
	return Model{
		config:  config,
		entries: entries,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Selected() *explorer.Entry {
	if m.entries == nil || len(*m.entries) == 0 {
		return nil
	}
	if m.cursor < 0 || m.cursor >= len(*m.entries) {
		return nil
	}
	return &(*m.entries)[m.cursor]
}

func (m Model) selectionCMD() tea.Cmd {
	entry := m.Selected()
	return func() tea.Msg {
		return events.SelectionChanged{Entry: *entry}
	}
}

func (m *Model) navigate() tea.Cmd {
	goToEntry := (*m.entries)[m.cursor]

	if goToEntry.Type != explorer.Dir {
		return nil
		// in the future this will "execute" the file
	}

	return func() tea.Msg {
		return events.WorkingDirectoryChanged{ChildName: goToEntry.Name}
	}
}

func (m *Model) goBack() tea.Cmd {
	return func() tea.Msg {
		return events.WorkingDirectoryChanged{ChildName: ".."}
	}
}

func (m *Model) calculateCursorFrame() {
	m.startCursor = 0
	m.endCursor = min(len(*m.entries), m.height-1)
}
