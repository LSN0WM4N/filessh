package events

import "github.com/LSN0WM4N/filessh/pkg/explorer"

type SetupExplorer struct{ Entries *[]explorer.Entry }

type SelectionChanged struct{ Entry explorer.Entry }
