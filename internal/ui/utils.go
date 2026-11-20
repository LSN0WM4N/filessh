package ui

import (
	"slices"

	"github.com/LSN0WM4N/filessh/internal/explorer"
)

func firstWithName(entries []explorer.Entry, name string) int {
	index := slices.IndexFunc(entries, func(e explorer.Entry) bool {
		return e.Name == name
	})

	return max(index, 0) // Not found is equal to set cursor to 0
}
