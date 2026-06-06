package explorer

import "time"

type Entry struct {
	Name  string
	IsDir bool
	Size  int64 // bytes

	// Unix file mode bits (e.g. 0755, 0644)
	Permissions uint16

	Mtime time.Time
}
