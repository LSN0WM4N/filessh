package ui

import (
	"cmp"

	"github.com/LSN0WM4N/filessh/pkg/explorer"
)

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
