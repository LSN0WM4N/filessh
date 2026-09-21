package explorer

type EntryType string

const (
	File     EntryType = "file"
	Dir      EntryType = "dir"
	Image    EntryType = "image"
	Audio    EntryType = "audio"
	Video    EntryType = "video"
	Document EntryType = "document"
	PDF      EntryType = "pdf"
)

type Entry struct {
	Name string
	Size int64 // IN bytes

	Type EntryType

	CTime int // Unix ctime
	MTime int // Unix mtime

	Permissions int16 // Unix-like permissions
}
