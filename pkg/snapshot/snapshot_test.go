package snapshot

import (
    "testing"
    "os"
    "path/filepath"
)

func TestCreateSnapshot(t *testing.T) {
    tmpDir := t.TempDir()
    testFile := filepath.Join(tmpDir, "test.txt")
    if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
        t.Fatal(err)
    }

    snap, err := CreateSnapshot(tmpDir)
    if err != nil {
        t.Fatal(err)
    }
    if len(snap.Files) != 1 {
        t.Errorf("expected 1 file, got %d", len(snap.Files))
    }
    if _, exists := snap.Files[testFile]; !exists {
        t.Errorf("expected file %s in snapshot", testFile)
    }
}