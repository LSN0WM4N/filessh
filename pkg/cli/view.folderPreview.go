package cli

import (
	"fmt"

	"github.com/LSN0WM4N/filessh/pkg/explorer"
)

func folderThumbnail(entry *explorer.Entry, width int) string {
	return fmt.Sprintf(
		"%c %s",
		FOLDER_ICON,
		truncateString(entry.Name, width),
	)
}
