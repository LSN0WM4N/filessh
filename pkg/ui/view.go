package ui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/LSN0WM4N/filessh/pkg/explorer"
	"github.com/LSN0WM4N/filessh/pkg/ui/format"
)

func (m model) View() tea.View {
	width := int(m.width)
	height := int(m.height)
	listHeight := height - BORDER_OFFSET

	if width <= 0 || height <= 0 || listHeight <= 0 {
		return tea.NewView("")
	}

	leftWidth, rightWidth := calculateColumnsWidths(width)
	entries, start, selected := m.computeVisibleWindow(listHeight)

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

	view := buildView(leftPanel, rightPanel)
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

		name := format.TruncateString(entry.Name, nameWidth)

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

func buildView(leftPanel string, rightPanel string) string {
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		leftPanel,
		strings.Repeat(" ", VIEW_GAP),
		rightPanel,
	)
}
