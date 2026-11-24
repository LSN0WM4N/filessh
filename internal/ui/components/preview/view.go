package preview

import (
	"fmt"

	"charm.land/lipgloss/v2"

	"github.com/LSN0WM4N/filessh/internal/explorer"
	"github.com/LSN0WM4N/filessh/internal/ui/format"
)

func (m Model) View() string {
	if m.entry == nil {
		return lipgloss.NewStyle().
			Width(m.width).
			Height(m.height).
			Render("")
	}

	var preview string

	switch m.entry.Type {
	case explorer.Document:
		preview = m.textFileThumbnail()
	case explorer.Audio:
		preview = m.audioFileThumbnail()
	case explorer.Image:
		preview = m.ImageFileThumbnail()
	case explorer.PDF:
		preview = m.PDFFileThumbnail()
	case explorer.Video:
		preview = m.VideoFileThumbnail()
	case explorer.Dir:
		preview = m.DirectoryFileThumbnail()
	default:
		preview = m.defaultFileThumbnail()
	}

	content := fmt.Sprintf(
		"%s\n\n%s\n%s\n%s",
		preview,
		m.entry.Name,
		format.UnixTime(m.entry.MTime),
		format.Size(m.entry.Size),
	)

	return lipgloss.NewStyle().
		Width(m.width).
		Height(m.height).
		MaxWidth(m.width).
		MaxHeight(m.height).
		Render(content)
}
