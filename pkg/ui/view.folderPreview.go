package ui

import (
	"fmt"

	"github.com/LSN0WM4N/filessh/pkg/explorer"
	"github.com/LSN0WM4N/filessh/pkg/ui/format"
)

func folderThumbnail(entry *explorer.Entry, width int) string {
	return fmt.Sprintf(
		"%c %s",
		FOLDER_ICON,
		format.TruncateString(entry.Name, width),
	)
}
