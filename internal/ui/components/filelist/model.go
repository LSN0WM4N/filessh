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
const LIST_CURSOR_MARGIN int = 3

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

func (m *Model) updateCursorView() {
	total := len(*m.entries)

	if total == 0 {
		m.startCursor = 0
		m.endCursor = 0
		m.cursor = 0
		return
	}

	m.cursor = max(0, min(m.cursor, total-1))

	viewHeight := max(m.height, 1)

	if m.cursor-LIST_CURSOR_MARGIN < m.startCursor {
		m.startCursor = m.cursor - LIST_CURSOR_MARGIN
	}

	if m.cursor+LIST_CURSOR_MARGIN >= m.endCursor {
		m.startCursor = m.cursor - viewHeight + 1 + LIST_CURSOR_MARGIN
	}

	m.startCursor = max(0, min(m.startCursor, max(total-viewHeight, 0)))
	m.endCursor = min(m.startCursor+viewHeight, total)
}

func (m *Model) isNotInCursorBounds(cursor int) bool {
	return m.startCursor > cursor || cursor > m.endCursor
}
