package explorer

import "os"

// ReadDir reads the contents of a directory and returns a list of entries
// (files and subdirectories) with their metadata.
//
// Parameters:
//   - path: A string representing the directory path to read.
//     Can be absolute or relative to the current working directory.
//
// Returns:
//   - []Entry: A slice of Entry structs containing name, type, size, and
//     modification time for each filesystem item in the directory.
//   - error: Returns an error if the directory cannot be accessed, does not
//     exist, or if the user lacks appropriate permissions.
//
// Behavior:
//   - The directory is read in the order returned by the underlying OS.
//   - If an individual entry's metadata cannot be retrieved (e.g., due to
//     permission issues or file system errors), that entry is silently
//     skipped and not included in the result.
//   - The function does not recurse into subdirectories.
//
// Example:
//
//	entries, err := ReadDir("/home/user/documents")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, entry := range entries {
//	    fmt.Printf("%s (%v, %d bytes)\n", entry.Name, entry.IsDir, entry.Size)
//	}
//
// Note: The returned entries are not sorted; consider using sort.Slice()
// if consistent ordering is required.
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
