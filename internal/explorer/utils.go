package explorer

import (
	"cmp"
	"path/filepath"
	"strings"

	"github.com/gabriel-vasile/mimetype"
)

func isValidNewPath(newPath, basePath string) bool {
	absoluteNewPath, err := filepath.Abs(newPath)
	if err != nil {
		return false
	}

	absoluteBasePath, err := filepath.Abs(basePath)
	if err != nil {
		return false
	}

	relativePath, err := filepath.Rel(absoluteBasePath, absoluteNewPath)
	if err != nil {
		return false
	}

	if relativePath == ".." ||
		strings.HasPrefix(relativePath, ".."+string(filepath.Separator)) {
		return false
	}

	return true
}

func shouldIgnore(ignoreHidden bool, filename string) bool {
	if !ignoreHidden || filename == ".." {
		return false
	} else {
		return filename[0] == '.'
	}
}

func getTypeFromMIME(path string) EntryType {
	mime, err := mimetype.DetectFile(path)
	if err != nil {
		return File
	}

	mimeString := mime.String()

	switch {
	case strings.HasPrefix(mimeString, "image/"):
		return Image

	case strings.HasPrefix(mimeString, "audio/"):
		return Audio

	case strings.HasPrefix(mimeString, "video/"):
		return Video

	case mimeString == "application/pdf":
		return PDF

	case strings.HasPrefix(mimeString, "text/"):
		return Document

	default:
		return File
	}
}

// compareByTypeFirst compares two Entry values first by type and then by name.
// Directories are considered smaller than files, so they appear first when
// sorting. If both have the same type, they are compared alphabetically
// by Name.
//
// It returns a negative value if a < b, a positive value if a > b, and zero
// if they are equal, so it is compatible with slices.SortFunc.
func compareByTypeFirst(a, b Entry) int {
	if a.Type == b.Type {
		return cmp.Compare(a.Name, b.Name)
	}
	if a.Type == Dir {
		return -1
	} else {
		return 1
	}
}
