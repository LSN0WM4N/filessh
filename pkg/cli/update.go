package cli

import tea "charm.land/bubbletea/v2"

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case entriesLoadedMsg:
		m.entries = msg.entries

	case tea.WindowSizeMsg:
		m.width = uint8(msg.Width & 0xFF)
		m.height = uint8(msg.Height & 0xFF)

	case tea.KeyPressMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.entries)-1 {
				m.cursor++
			}

		case "space", "enter":
			newDirName := m.entries[m.cursor].Name

			reloadExplorerWithPath(&m, newDirName)

		case "backspace":
			reloadExplorerWithPath(&m, "..")

		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, nil
}

func reloadExplorerWithPath(m *model, newDirName string) {
	success, _ := m.explorer.ChangeDir(newDirName)

	if !success {
		return
	}

	m.entries, _ = m.explorer.List()
	m.cursor = 0
}
