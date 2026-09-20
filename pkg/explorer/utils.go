package explorer

import (
	"path/filepath"
	"strings"
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
