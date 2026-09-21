package cli

import (
	"cmp"

	"github.com/LSN0WM4N/filessh/pkg/explorer"
)

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
