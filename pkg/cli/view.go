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

	const gap = 2

	leftWidth := int(float64(width) * 0.55)
	rightWidth := width - leftWidth - gap

	if leftWidth < 10 {
		leftWidth = 10
	}

	if rightWidth < 10 {
		rightWidth = 10
	}

	listHeight := height - BORDER_OFFSET

	if listHeight <= 0 {
		return tea.NewView("")
	}

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

	var selected *explorer.Entry

	if m.cursor >= 0 && m.cursor < len(m.entries) {
		selected = &m.entries[m.cursor]
	}

	leftPanel := buildLeftPanel(
		start,
		m.cursor,
		leftWidth,
		listHeight,
		&entries,
	)

	rightPanel := buildRightPanel(
		selected,
		rightWidth,
		listHeight,
	)

	view := lipgloss.JoinHorizontal(
		lipgloss.Top,
		leftPanel,
		strings.Repeat(" ", gap),
		rightPanel,
	)

	return tea.NewView(view)
}

func buildLeftPanel(
	start int,
	modelCursor int,
	leftWidth int,
	listHeight int,
	entries *[]explorer.Entry,
) string {

	var leftLines []string

	for i, entry := range *entries {
		cursor := " "
		icon := CharacterToDIsplay("FOLDER")

		absoluteIndex := start + i

		if absoluteIndex == modelCursor {
			cursor = ">"
		}

		if entry.Type != explorer.Dir {
			icon = CharacterToDIsplay(strings.ToUpper(string(entry.Type)))
		}

		if entry.Name == ".." {
			icon = CharacterToDIsplay("BACK")
		}

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

	return lipgloss.NewStyle().
		Width(leftWidth).
		Height(listHeight).
		MaxWidth(leftWidth).
		MaxHeight(listHeight).
		Render(leftContent)
}

func buildRightPanel(
	selected *explorer.Entry,
	rightWidth int,
	listHeight int,
) string {
	rightContent := renderPreview(
		selected,
		rightWidth,
		listHeight,
	)

	return lipgloss.NewStyle().
		Width(rightWidth).
		Height(listHeight).
		MaxWidth(rightWidth).
		MaxHeight(listHeight).
		Render(rightContent)
}
