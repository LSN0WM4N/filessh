package cli

import tea "charm.land/bubbletea/v2"

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case entriesLoadedMsg:
		m.entries = msg.entries

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
			m.explorer.ChangeDir(newDirName)
			m.entries, _ = m.explorer.List()
			m.cursor = 0

		case "backspace":
			m.explorer.ChangeDir("..")
			m.entries, _ = m.explorer.List()
			m.cursor = 0

		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, nil
}
