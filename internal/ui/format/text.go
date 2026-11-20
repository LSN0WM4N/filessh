package format

import "charm.land/lipgloss/v2"

// truncateString shortens s so that its visible width does not exceed
// maxWidth, appending "..." at the end when truncation occurs.
//
// It returns an empty string if maxWidth <= 0. If maxWidth <= 3, no ellipsis
// is added and the string is simply cut.
func TruncateString(s string, maxWidth int) string {
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
