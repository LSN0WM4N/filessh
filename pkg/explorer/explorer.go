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

func (e *Explorer) List() ([]Entry, error) {
	results := make([]Entry, 0)

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

		entryType := File
		if entry.IsDir() {
			entryType = Dir
		}

		if entry.Name()[0] == '.' {
			continue // ignore hidden files by now
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
