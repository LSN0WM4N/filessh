package explorer

import "os"

func ReadDir(path string) ([]Entry, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var result []Entry
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue
		}
		result = append(result, Entry{
			Name:  entry.Name(),
			IsDir: entry.IsDir(),
			Size:  info.Size(),
			Mtime: info.ModTime(),
		})
	}
	return result, nil
}
