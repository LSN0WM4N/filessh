package explorer

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExplorerList(t *testing.T) {
	tempDir := t.TempDir()

	file1 := filepath.Join(tempDir, "file1.txt")
	file2 := filepath.Join(tempDir, "file2.txt")

	if err := os.WriteFile(file1, []byte("hello"), 0644); err != nil {
		t.Fatal(err)
	}

	if err := os.WriteFile(file2, []byte("hello world"), 0644); err != nil {
		t.Fatal(err)
	}

	dir1 := filepath.Join(tempDir, "dir1")
	dir2 := filepath.Join(tempDir, "dir2")

	if err := os.Mkdir(dir1, 0755); err != nil {
		t.Fatal(err)
	}

	if err := os.Mkdir(dir2, 0755); err != nil {
		t.Fatal(err)
	}

	explorer := New(tempDir)

	entries, err := explorer.List(false, false)
	if err != nil {
		t.Fatalf("List() returned an error: %v", err)
	}

	if len(entries) != 4 {
		t.Fatalf("expected 4 entries, got %d", len(entries))
	}

	expected := map[string]bool{
		"file1.txt": false,
		"file2.txt": false,
		"dir1":      false,
		"dir2":      false,
	}

	for _, entry := range entries {
		if _, exists := expected[entry.Name]; !exists {
			t.Errorf("unexpected entry: %s", entry.Name)
			continue
		}

		expected[entry.Name] = true

		switch entry.Name {
		case "file1.txt", "file2.txt":
			if entry.Type != File {
				t.Errorf(
					"expected %s to be a file, got %s",
					entry.Name,
					entry.Type,
				)
			}

		case "dir1", "dir2":
			if entry.Type != Dir {
				t.Errorf(
					"expected %s to be a directory, got %s",
					entry.Name,
					entry.Type,
				)
			}
		}
	}

	for name, found := range expected {
		if !found {
			t.Errorf("expected entry %s was not found", name)
		}
	}
}

func TestExplorerListEmptyDirectory(t *testing.T) {
	tempDir := t.TempDir()

	explorer := New(tempDir)

	entries, err := explorer.List(false, false)
	if err != nil {
		t.Fatalf("List() returned an error: %v", err)
	}

	if len(entries) != 0 {
		t.Fatalf("expected empty list, got %d entries", len(entries))
	}
}

func TestExplorerListNonExistingDirectory(t *testing.T) {
	tempDir := t.TempDir()

	nonExistingPath := filepath.Join(tempDir, "does-not-exist")

	explorer := New(nonExistingPath)

	entries, err := explorer.List(false, false)

	if err == nil {
		t.Fatal("expected an error, got nil")
	}

	if entries == nil {
		t.Fatal("expected non-nil entries slice")
	}

	if len(entries) != 0 {
		t.Fatalf("expected empty entries, got %d", len(entries))
	}
}

func TestExplorerChangeDir(t *testing.T) {
	tempDir := t.TempDir()

	childDir := filepath.Join(tempDir, "child")

	if err := os.Mkdir(childDir, 0755); err != nil {
		t.Fatal(err)
	}

	explorer := New(tempDir)

	changed, err := explorer.ChangeDir("child")

	if err != nil {
		t.Fatalf("ChangeDir() returned an error: %v", err)
	}

	if !changed {
		t.Fatal("expected ChangeDir() to return true")
	}

	expectedPath, err := filepath.Abs(childDir)
	if err != nil {
		t.Fatal(err)
	}

	if explorer.path != expectedPath {
		t.Fatalf(
			"expected path %q, got %q",
			expectedPath,
			explorer.path,
		)
	}
}

func TestExplorerChangeDirNonExistingDirectory(t *testing.T) {
	tempDir := t.TempDir()

	explorer := New(tempDir)

	changed, err := explorer.ChangeDir("does-not-exist")

	if err == nil {
		t.Fatal("expected an error, got nil")
	}

	if changed {
		t.Fatal("expected ChangeDir() to return false")
	}

	if explorer.path != tempDir {
		t.Fatalf(
			"Path should not change, expected %q, got %q",
			tempDir,
			explorer.path,
		)
	}
}

func TestExplorerChangeDirFile(t *testing.T) {
	tempDir := t.TempDir()

	filePath := filepath.Join(tempDir, "file.txt")

	if err := os.WriteFile(filePath, []byte("hello"), 0644); err != nil {
		t.Fatal(err)
	}

	explorer := New(tempDir)

	changed, err := explorer.ChangeDir("file.txt")

	if err != nil {
		t.Fatalf("ChangeDir() returned an error: %v", err)
	}

	if changed {
		t.Fatal("expected ChangeDir() to return false for a file")
	}

	if explorer.path != tempDir {
		t.Fatalf(
			"Path should not change, expected %q, got %q",
			tempDir,
			explorer.path,
		)
	}
}

func TestExplorerChangeDirParent(t *testing.T) {
	tempDir := t.TempDir()

	childDir := filepath.Join(tempDir, "child")

	if err := os.Mkdir(childDir, 0755); err != nil {
		t.Fatal(err)
	}

	explorer := New(tempDir)

	changed, err := explorer.ChangeDir("..")

	if err != nil {
		t.Fatalf("ChangeDir() returned an error: %v", err)
	}

	if changed {
		t.Fatal("expected ChangeDir(\"..\") to return false")
	}

	if explorer.path != tempDir {
		t.Fatalf(
			"Path should not change, expected %q, got %q",
			tempDir,
			explorer.path,
		)
	}
}

func TestExplorerChangeDirParentFromChild(t *testing.T) {
	tempDir := t.TempDir()

	childDir := filepath.Join(tempDir, "child")

	if err := os.Mkdir(childDir, 0755); err != nil {
		t.Fatal(err)
	}

	explorer := New(tempDir)

	changed, err := explorer.ChangeDir("child")

	if err != nil {
		t.Fatalf("failed to enter child: %v", err)
	}

	if !changed {
		t.Fatal("expected to enter child directory")
	}

	changed, err = explorer.ChangeDir("..")

	if err != nil {
		t.Fatalf("failed to go to parent: %v", err)
	}

	if !changed {
		t.Fatal("expected to go back to parent")
	}

	if explorer.path != tempDir {
		t.Fatalf(
			"expected path %q, got %q",
			tempDir,
			explorer.path,
		)
	}
}

func TestExplorerChangeDirCannotEscapeBasePath(t *testing.T) {
	tempDir := t.TempDir()

	childDir := filepath.Join(tempDir, "child")

	if err := os.Mkdir(childDir, 0755); err != nil {
		t.Fatal(err)
	}

	explorer := New(tempDir)

	changed, err := explorer.ChangeDir("child")

	if err != nil {
		t.Fatal(err)
	}

	if !changed {
		t.Fatal("expected to enter child")
	}

	changed, err = explorer.ChangeDir("../..")

	if err != nil {
		t.Fatalf("ChangeDir() returned an error: %v", err)
	}

	if changed {
		t.Fatal("expected ChangeDir() to prevent escaping basePath")
	}

	expectedPath, err := filepath.Abs(childDir)
	if err != nil {
		t.Fatal(err)
	}

	if explorer.path != expectedPath {
		t.Fatalf(
			"expected path to remain %q, got %q",
			expectedPath,
			explorer.path,
		)
	}
}
