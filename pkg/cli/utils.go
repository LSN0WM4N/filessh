package cli

import (
	"cmp"
	"fmt"
	"os"
	"strings"
	"time"

	"charm.land/lipgloss/v2"
	"github.com/LSN0WM4N/filessh/pkg/explorer"
)

var supportUTF8Result *bool
var forceBackwardFlag bool

func SetForceBackward(newValue bool) {
	forceBackwardFlag = newValue
}

func compareByTypeFirst(a, b explorer.Entry) int {
	if a.Type == b.Type {
		return cmp.Compare(a.Name, b.Name)
	}
	if a.Type == explorer.Dir {
		return -1
	} else {
		return 1
	}
}

func renderImagePlaceholder(width int) string {
	if !supportsUTF8() {
		return ""
	}

	imageWidth := 18

	if width < imageWidth {
		imageWidth = width
	}

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

func formatUnixTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("02/01/2006 15:04")
}

func formatSize(size int64) string {
	const unit = 1024

	if size < unit {
		return fmt.Sprintf("%d B", size)
	}

	div, exp := int64(unit), 0

	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	return fmt.Sprintf(
		"%.1f %cB",
		float64(size)/float64(div),
		"KMGTPE"[exp],
	)
}

func truncateString(s string, maxWidth int) string {
	if maxWidth <= 0 {
		return ""
	}

	if lipgloss.Width(s) <= maxWidth {
		return s
	}

	if maxWidth <= 3 {
		return s[:maxWidth]
	}

	runes := []rune(s)

	if len(runes) <= maxWidth {
		return s
	}

	return string(runes[:maxWidth-3]) + "..."
}

func CharacterToDIsplay(name string) rune {
	realName := name
	if !supportsUTF8() {
		realName += "_RETRO"
	}

	switch realName {
	case "FOLDER":
		return FOLDER_ICON
	case "FILE":
		return FILE_ICON
	case "BACK":
		return BACK_ICON

	case "FOLDER_RETRO":
		return FOLDER_ICON_RETRO
	case "FILE_RETRO":
		return FILE_ICON_RETRO
	case "BACK_RETRO":
		return BACK_ICON_RETRO

	default:
		return ' '
	}
}

// EXPERIMENTAL
// TODO: Just ask one time for performance
func supportsUTF8() bool {
	if forceBackwardFlag {
		return false
	}

	if supportUTF8Result != nil {
		return *supportUTF8Result
	}

	result := false
	for _, key := range []string{"LC_ALL", "LC_CTYPE", "LANG"} {
		value := os.Getenv(key)

		if strings.Contains(strings.ToLower(value), "utf-8") ||
			strings.Contains(strings.ToLower(value), "utf8") {
			result = true
		}
	}

	supportUTF8Result = &result
	return result
}
