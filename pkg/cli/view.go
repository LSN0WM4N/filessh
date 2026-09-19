package cli

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

func (m model) View() tea.View {
	var s string
	entries := m.entries

	for i, entry := range entries {
		cursor := " "
		if i == m.cursor {
			cursor = ">"
		}

		s += fmt.Sprintf("%s [%s]\n", cursor, entry.Name)
	}

	return tea.NewView(s)
}
