package explorer

import (
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
