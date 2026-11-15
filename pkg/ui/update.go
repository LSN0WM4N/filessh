package ui

import (
	tea "charm.land/bubbletea/v2"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	m.filelist, cmd = m.filelist.Update(msg)
	cmds = append(cmds, cmd)

	m.preview, cmd = m.preview.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

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

	return m, tea.Batch(cmds...)
}
