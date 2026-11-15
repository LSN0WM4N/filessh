package ui

import (
	tea "charm.land/bubbletea/v2"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.filelist.Update(msg)
	m.preview.Update(msg)

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = uint8(msg.Width & 0xFF)
		m.height = uint8(msg.Height & 0xFF)

	case tea.KeyPressMsg:
		switch msg.String() {
		case "up", "k":

		case "down", "j":

		case "space", "enter":

		case "backspace":

		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, nil
}
