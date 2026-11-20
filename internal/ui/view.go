package ui

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (m Model) View() tea.View {
	width := int(m.width)
	height := int(m.height)
	listHeight := height - BORDER_OFFSET

	if width <= 0 || height <= 0 || listHeight <= 0 {
		return tea.NewView("")
	}

	view := lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.filelist.View(),
		strings.Repeat(" ", VIEW_GAP),
		m.preview.View(),
	)
	return tea.NewView(view)
}
