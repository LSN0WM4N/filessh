package filelist

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/LSN0WM4N/filessh/internal/explorer"
	"github.com/LSN0WM4N/filessh/internal/ui/format"
	"github.com/LSN0WM4N/filessh/internal/ui/theme"
)

func (m Model) View() string {
	var leftLines []string

	for i, entry := range *m.entries {
		cursor := " "
		icon := theme.CharacterToDIsplay(m.config, "FOLDER")

		absoluteIndex := m.startCursor + i

		if absoluteIndex == m.cursor {
			cursor = ">"
		}

		// FIXIT
		if entry.Type != explorer.Dir {
			icon = theme.CharacterToDIsplay(m.config, strings.ToUpper(string(entry.Type)))
		}

		if entry.Name == ".." {
			icon = theme.CharacterToDIsplay(m.config, "BACK")
		}

		nameWidth := m.width - 5

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
		Width(m.width).
		Height(m.height).
		MaxWidth(m.width).
		MaxHeight(m.height).
		Render(leftContent)
}
