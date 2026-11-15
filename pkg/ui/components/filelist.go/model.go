package filelist

import (
	tea "charm.land/bubbletea/v2"
	"github.com/LSN0WM4N/filessh/pkg/events"
	"github.com/LSN0WM4N/filessh/pkg/explorer"
	"github.com/LSN0WM4N/filessh/pkg/ui/theme"
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

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case events.SetupExplorer:
		m.entries = msg.Entries
		m.cursor = 0
		m.startCursor = 0

	case tea.WindowSizeMsg:
		m.width = int(float32(msg.Width) * VIEW_GAP)
		m.height = int(float32(msg.Height) * VIEW_GAP)
		m.endCursor = min(len(*m.entries), m.height)

	case tea.KeyPressMsg:
		switch msg.String() {
		case "up", "k":
			m.cursor = max(m.cursor-1, 0)

		case "down", "j":
			m.cursor = min(m.cursor+1, m.endCursor)

		case "space", "enter":

		case "backspace":

		case "q", "ctrl+c":
			return nil, tea.Quit
		}
	}

	return nil, nil
}
