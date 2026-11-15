package preview

import (
	"strings"
)

func (m *Model) textFileThumbnail() string {
	return "TEXT"
}

func (m *Model) audioFileThumbnail() string {
	return "AUDIO"
}

func (m *Model) defaultFileThumbnail() string {
	if !m.config.UFT8() {
		return ""
	}

	imageWidth := min(m.width, 18)

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
