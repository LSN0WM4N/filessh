package explorer

import (
	"os"
	"path/filepath"
)

type Explorer struct {
	path string

	basePath string
}

func (e *Explorer) Path() string {
	return e.path
}

// A function that list all the subdirectories in the explorer directory
// it takes `includeBackwardPath` for include de '..' path to go back
// and `ignoreHidden` for not include files or directories that start with '.'
func (e *Explorer) List(includeBackwardPath, ignoreHidden bool) ([]Entry, error) {
	results := make([]Entry, 0)

	if includeBackwardPath {
		results = append(results, Entry{
			Name: "..",
			Size: 4096,

			Type: Dir,

			CTime:       0,
			MTime:       0,
			Permissions: 0xFF,
		})
	}

	rawList, err := os.ReadDir(e.path)
	if err != nil {
		// At this point, results is an empty array
		return results, err
	}

	for _, entry := range rawList {
		info, err := entry.Info()
		if err != nil {
			return results, err
		}

		entryType := Dir
		if !entry.IsDir() {
			entryType = getTypeFromMIME(filepath.Join(e.path, entry.Name()))
		}

		if shouldIgnore(ignoreHidden, entry.Name()) {
			continue
		}

		results = append(results, Entry{
			Name:        entry.Name(),
			Size:        info.Size(),
			Type:        entryType,
			CTime:       int(info.ModTime().Unix()),
			MTime:       int(info.ModTime().Unix()),
			Permissions: int16(info.Mode().Perm()),
		})
	}

	return results, nil
}

func (e *Explorer) ChangeDir(dirName string) (bool, error) {
	newPath := filepath.Join(e.path, dirName)

	if !isValidNewPath(newPath, e.basePath) {
		return false, nil
	}

	info, err := os.Stat(newPath)

	if err != nil || !info.IsDir() {
		return false, err
	}

	e.path = newPath

	return true, nil
}

func New(path string) *Explorer {
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		return nil
	}

	newExplorer := Explorer{
		path:     absolutePath,
		basePath: absolutePath, // !TODO: Smarter way to define a base path
	}

	return &newExplorer
}
