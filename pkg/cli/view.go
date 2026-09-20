package cli

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

func (m model) View() tea.View {
	var s string

	entries := m.entries

	start := 0
	end := len(m.entries)

	height := int(m.height)&0xFF - BORDER_OFFSET

	if height <= 0 {
		return tea.NewView("")
	}

	if len(m.entries) > height {
		cursorOffset := CURSOR_OFFSET

		end = height

		if m.cursor > cursorOffset {
			start = m.cursor - cursorOffset
			end = start + height
		}

		if end > len(m.entries) {
			end = len(m.entries)
			start = end - height
		}

		entries = m.entries[start:end]
	}

	for i, entry := range entries {
		cursor := " "

		if i == m.cursor-start {
			cursor = ">"
		}

		s += fmt.Sprintf("%s [%s]\n", cursor, entry.Name)
	}

	return tea.NewView(s)
}
