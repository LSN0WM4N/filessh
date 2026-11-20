package filelist

import (
	"log"

	tea "charm.land/bubbletea/v2"

	"github.com/LSN0WM4N/filessh/internal/events"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case events.SetupExplorer:
		m.entries = msg.Entries
		m.cursor = 0

	case tea.WindowSizeMsg:
		log.Printf("[ui.filelist.Model.Update:WindowSizeMsg] Recived { Width: %d, Height: %d }", msg.Width, msg.Height)
		m.width = int(float32(msg.Width) * VIEW_GAP)
		m.height = msg.Height

		m.calculateCursorFrame()

	case tea.KeyPressMsg:
		switch msg.String() {
		// TOFIX: Bug navegating when the windows is smallet than the total files
		case "up", "k":
			m.cursor = max(m.cursor-1, 0)
			return m, m.selectionCMD()

		case "down", "j":
			m.cursor = min(m.cursor+1, m.endCursor)
			return m, m.selectionCMD()

		case "space", "enter":
			return m, m.navigate()

		case "backspace":
			return m, m.goBack()

		case "q", "ctrl+c":
			return m, tea.Quit
		}

	case events.DirectoryListChanged:
		m.entries = msg.DirectoryList

		m.cursor = 0
		m.calculateCursorFrame()
	}

	return m, nil
}
