package cli

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/LSN0WM4N/filessh/pkg/explorer"
)

func (m model) View() tea.View {
	width := int(m.width)
	height := int(m.height)

	if width <= 0 || height <= 0 {
		return tea.NewView("")
	}

	// Reservamos una pequeña separación entre los paneles.
	const gap = 2

	// El panel izquierdo ocupa el 55%.
	leftWidth := int(float64(width) * 0.55)
	rightWidth := width - leftWidth - gap

	if leftWidth < 10 {
		leftWidth = 10
	}

	if rightWidth < 10 {
		rightWidth = 10
	}

	// Altura disponible para los entries.
	listHeight := height - BORDER_OFFSET

	if listHeight <= 0 {
		return tea.NewView("")
	}

	// ---------------------------------------------------------
	// WINDOW DE ENTRIES
	// ---------------------------------------------------------

	start := 0
	end := len(m.entries)

	if len(m.entries) > listHeight {
		end = listHeight

		if m.cursor > CURSOR_OFFSET {
			start = m.cursor - CURSOR_OFFSET
			end = start + listHeight
		}

		if end > len(m.entries) {
			end = len(m.entries)
			start = end - listHeight
		}
	}

	entries := m.entries[start:end]

	// ---------------------------------------------------------
	// PANEL IZQUIERDO
	// ---------------------------------------------------------

	var leftLines []string

	for i, entry := range entries {
		cursor := " "
		icon := FOLDER_ICON

		absoluteIndex := start + i

		if absoluteIndex == m.cursor {
			cursor = ">"
		}

		if entry.Type == explorer.File {
			icon = FILE_ICON
		}

		if entry.Name == ".." {
			icon = BACK_ICON
		}

		// Espacio disponible para el nombre.
		//
		// cursor + espacio + icon + espacio = 5 caracteres
		nameWidth := leftWidth - 5

		if nameWidth < 1 {
			nameWidth = 1
		}

		name := truncateString(entry.Name, nameWidth)

		line := fmt.Sprintf(
			"%s %c %s",
			cursor,
			icon,
			name,
		)

		leftLines = append(leftLines, line)
	}

	leftContent := strings.Join(leftLines, "\n")

	leftPanel := lipgloss.NewStyle().
		Width(leftWidth).
		Height(listHeight).
		MaxWidth(leftWidth).
		MaxHeight(listHeight).
		Render(leftContent)

	// ---------------------------------------------------------
	// PANEL DERECHO
	// ---------------------------------------------------------

	var selected *explorer.Entry

	if m.cursor >= 0 && m.cursor < len(m.entries) {
		selected = &m.entries[m.cursor]
	}

	rightContent := renderPreview(
		selected,
		rightWidth,
		listHeight,
	)

	rightPanel := lipgloss.NewStyle().
		Width(rightWidth).
		Height(listHeight).
		MaxWidth(rightWidth).
		MaxHeight(listHeight).
		Render(rightContent)

	// ---------------------------------------------------------
	// LAYOUT FINAL
	// ---------------------------------------------------------

	view := lipgloss.JoinHorizontal(
		lipgloss.Top,
		leftPanel,
		strings.Repeat(" ", gap),
		rightPanel,
	)

	return tea.NewView(view)
}
