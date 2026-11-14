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
	contentWidth := max(width-4, 10)

	if entry.Type == explorer.Dir {
		return folderThumbnail(entry, contentWidth)
	}

	name := entry.Name

	preview := fileThumbnail(entry, contentWidth)

	return fmt.Sprintf(
		"%s\n\n%s\n%s\n%s",
		preview,
		name,
		formatUnixTime(entry.MTime),
		formatSize(entry.Size),
	)
}
