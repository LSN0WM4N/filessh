package preview

import (
	tea "charm.land/bubbletea/v2"
	"github.com/LSN0WM4N/filessh/pkg/events"
	"github.com/LSN0WM4N/filessh/pkg/explorer"
	"github.com/LSN0WM4N/filessh/pkg/ui/theme"
)

type Model struct {
	config *theme.Config
	entry  *explorer.Entry

	width  int
	height int
}

const VIEW_GAP float32 = 0.45

func New(config *theme.Config, entry *explorer.Entry) Model {
	return Model{
		config: config,
		entry:  entry,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

// Selection change will recive an Entry with te new entry to be displayed
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case events.SelectionChanged:
		entry := msg.Entry
		m.entry = &entry

	case tea.WindowSizeMsg:
		m.width = int(float32(msg.Width) * VIEW_GAP)
		m.height = int(float32(msg.Height) * VIEW_GAP)
	}

	return nil, nil
}
