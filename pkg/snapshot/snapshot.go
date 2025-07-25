package snapshot

import (
    "crypto/sha256"
    "encoding/hex"
    "io"
    "os"
    "path/filepath"
    "time"
)

type Snapshot struct {
    ID        string
    Timestamp time.Time
    Files     map[string]string
}

func CreateSnapshot(dir string) (*Snapshot, error) {
    snap := &Snapshot{
        ID:        generateID(),
        Timestamp: time.Now(),
        Files:     make(map[string]string),
    }

    err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() {
            hash, err := computeFileHash(path)
            if err != nil {
                return err
            }
            snap.Files[path] = hash
        }
        return nil
    })
    if err != nil {
        return nil, err
    }
    return snap, nil
}

func generateID() string {
    hasher := sha256.New()
    hasher.Write([]byte(time.Now().String()))
    return hex.EncodeToString(hasher.Sum(nil))
}

func computeFileHash(path string) (string, error) {
    file, err := os.Open(path)
    if err != nil {
        return "", err
    }
    defer file.Close()

    hasher := sha256.New()
    if _, err := io.Copy(hasher, file); err != nil {
        return "", err
    }
    return hex.EncodeToString(hasher.Sum(nil)), nil
}