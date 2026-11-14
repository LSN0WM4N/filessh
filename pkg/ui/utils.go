package ui

import (
	"cmp"
	"os"
	"strings"

	"github.com/LSN0WM4N/filessh/pkg/explorer"
)

// supportUTF8Result caches the result of the terminal UTF-8 support
// detection. It is nil until supportsUTF8 runs for the first time.
var supportUTF8Result *bool

// forceBackwardFlag forces compatibility mode (retro icons without UTF-8)
// regardless of the environment configuration.
var forceBackwardFlag bool

// SetForceBackwardCompatibility enables or disables backward compatibility mode.
// When newValue is true, the interface always uses the "retro" ASCII
// characters instead of UTF-8 icons, regardless of the terminal configuration.
func SetForceBackwardCompatibility(newValue bool) {
	forceBackwardFlag = newValue
}

// compareByTypeFirst compares two Entry values first by type and then by name.
// Directories are considered smaller than files, so they appear first when
// sorting. If both have the same type, they are compared alphabetically
// by Name.
//
// It returns a negative value if a < b, a positive value if a > b, and zero
// if they are equal, so it is compatible with slices.SortFunc.
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

// supportsUTF8 reports whether the terminal supports UTF-8 by inspecting the
// LC_ALL, LC_CTYPE and LANG environment variables for "utf-8" or "utf8"
// (case-insensitive).
//
// It always returns false while forceBackwardFlag is set. The result is
// cached in supportUTF8Result, so detection is only performed once.
//
// EXPERIMENTAL: the environment-variable heuristic may not be reliable on
// every terminal and its behavior may change.
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

// calculateColumnsWidths splits the total width between the left and right
// columns of the view. The left column takes the LEFT_WIDTH_PERCENTAGE
// proportion of the total and the right column takes the rest, minus the
// VIEW_GAP separation space. Both columns have a minimum width of 10.
//
// It returns (leftWidth, rightWidth).
func calculateColumnsWidths(width int) (int, int) {
	leftWidth := int(float64(width) * LEFT_WIDTH_PERCENTAGE)
	rightWidth := width - leftWidth - VIEW_GAP

	return max(leftWidth, 10), max(rightWidth, 10)
}

// CharacterToDIsplay returns the icon (rune) associated with the logical name
// of an item type: "FOLDER", "FILE", "IMAGE", "AUDIO", "VIDEO", "PDF",
// "DOCUMENT" or "BACK".
//
// If the terminal supports UTF-8, Unicode icons are returned; otherwise, the
// ASCII-compatible "retro" variants are returned. In retro mode, all file
// types share the same file icon (FILE_ICON_RETRO). If the name is not
// recognized, it returns a space.
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
	case "IMAGE":
		return IMAGE_ICON
	case "AUDIO":
		return AUDIO_ICON
	case "VIDEO":
		return VIDEO_ICON
	case "PDF":
		fallthrough
	case "DOCUMENT":
		return PDF_ICON
	case "BACK":
		return BACK_ICON

	case "FOLDER_RETRO":
		return FOLDER_ICON_RETRO
	case "FILE_RETRO":
		fallthrough
	case "IMAGE_RETRO":
		fallthrough
	case "AUDIO_RETRO":
		fallthrough
	case "VIDEO_RETRO":
		fallthrough
	case "PDF_RETRO":
		fallthrough
	case "DOCUMENT_RETRO":
		return FILE_ICON_RETRO
	case "BACK_RETRO":
		return BACK_ICON_RETRO

	default:
		return ' '
	}
}
