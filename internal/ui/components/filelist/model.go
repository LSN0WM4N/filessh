package filelist

import (
	"log"

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

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	log.Printf("[ui.filelist.Model.Update] Recived message type: %T", msg)

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
			return m, m.selectionCMD()

		case "down", "j":
			m.cursor = min(m.cursor+1, m.endCursor)
			return m, m.selectionCMD()

		case "space", "enter":
			return m, m.navigate()

		case "backspace":

		case "q", "ctrl+c":
			return m, tea.Quit
		}

	case events.DirectoryListChanged:
		log.Print("Payload: ")
		log.Print(*msg.DirectoryList)

		m.entries = msg.DirectoryList

		// TODO: Create method for setting up the cursor bounds
		m.cursor = 0
		m.startCursor = 0
		m.endCursor = len(*m.entries)
	}

	return m, nil
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

func (m Model) navigate() tea.Cmd {
	goToEntry := (*m.entries)[m.cursor]

	if goToEntry.Type != explorer.Dir {
		return nil
		// in the future this will "execute" the file
	}

	return func() tea.Msg {
		return events.WorkingDirectoryChanged{ChildName: goToEntry.Name}
	}
}
