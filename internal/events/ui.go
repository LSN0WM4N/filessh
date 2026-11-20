package events

import "github.com/LSN0WM4N/filessh/internal/explorer"

type SetupExplorer struct{ Entries *[]explorer.Entry }
type SelectionChanged struct{ Entry explorer.Entry }
type WorkingDirectoryChanged struct{ ChildName string }

type DirectoryListChanged struct {
	DirectoryList  *[]explorer.Entry
	CursorPosition int
}
