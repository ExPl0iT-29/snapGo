package dedup

import (
    "os"
    "path/filepath"
)

func IsDuplicate(hash string, storageDir string) bool {
    _, err := os.Stat(filepath.Join(storageDir, hash[:2], hash[2:]))
    return err == nil
}