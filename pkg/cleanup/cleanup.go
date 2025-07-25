package cleanup

import (
    "os"
    "time"
    "path/filepath"
)

func CleanupSnapshots(snapshotDir string, maxAge time.Duration) error {
    return filepath.Walk(snapshotDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() && time.Since(info.ModTime()) > maxAge {
            return os.Remove(path)
        }
        return nil
    })
}