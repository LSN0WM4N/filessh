package cli

import (
	"strings"

	"github.com/LSN0WM4N/filessh/pkg/explorer"
)

func fileThumbnail(entry *explorer.Entry, width int) string {
	switch entry.Type {
	case explorer.Document:
		return textFileThumbnail(width)
	case explorer.Audio:
		return audioFileThumbnail(width)
	default:
		return defaultFileThumbnail(width)
	}
}

func textFileThumbnail(width int) string {
	return "TEXT"
}

func audioFileThumbnail(width int) string {
	return "AUDIO"
}

func defaultFileThumbnail(width int) string {
	if !supportsUTF8() {
		return ""
	}

	imageWidth := min(width, 18)

	if imageWidth < 4 {
		return ""
	}

	imageHeight := 8

	border := "┌" + strings.Repeat("─", imageWidth-2) + "┐"
	bottom := "└" + strings.Repeat("─", imageWidth-2) + "┘"

	var s strings.Builder

	s.WriteString(border)
	s.WriteByte('\n')

	for i := 0; i < imageHeight-2; i++ {
		insideWidth := imageWidth - 2

		text := ""

		if i == (imageHeight-2)/2 {
			text = "PREVIEW"
		}

		paddingLeft := (insideWidth - len(text)) / 2
		paddingRight := insideWidth - paddingLeft - len(text)

		line := "│"
		line += strings.Repeat(" ", paddingLeft)
		line += text
		line += strings.Repeat(" ", paddingRight)
		line += "│"

		s.WriteString(line)
		s.WriteByte('\n')
	}

	s.WriteString(bottom)

	return s.String()
}
