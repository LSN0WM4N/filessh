package cli

import (
	"fmt"

	"charm.land/lipgloss/v2"
	"github.com/LSN0WM4N/filessh/pkg/explorer"
)

var (
	leftPanel = lipgloss.NewStyle().
			Width(50).
			PaddingRight(2)

	rightPanel = lipgloss.NewStyle().
			Width(30).
			PaddingLeft(2)

	previewStyle = lipgloss.NewStyle().
			Align(lipgloss.Center)
)

func renderPreview(entry *explorer.Entry, width int, height int) string {
	if entry == nil {
		return ""
	}

	if entry.Type == explorer.Dir {
		return fmt.Sprintf(
			"%c %s",
			FOLDER_ICON,
			truncateString(entry.Name, width-4),
		)
	}

	contentWidth := width - 4

	if contentWidth < 10 {
		contentWidth = 10
	}

	name := entry.Name

	preview := renderImagePlaceholder(contentWidth)

	return fmt.Sprintf(
		"%s\n\n%s\n%s\n%s",
		preview,
		name,
		formatUnixTime(entry.MTime),
		formatSize(entry.Size),
	)
}
