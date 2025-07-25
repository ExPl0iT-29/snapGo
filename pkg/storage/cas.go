package storage

import (
    "crypto/sha256"
    "fmt"
    "io"
    "os"
    "path/filepath"
)

func StoreFile(filePath string, storageDir string) (string, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return "", err
    }
    defer file.Close()

    hasher := sha256.New()
    if _, err := io.Copy(hasher, file); err != nil {
        return "", err
    }
    hash := fmt.Sprintf("%x", hasher.Sum(nil))

    destPath := filepath.Join(storageDir, hash[:2], hash[2:])
    if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
        return "", err
    }
    if err := copyFile(filePath, destPath); err != nil {
        return "", err
    }
    return hash, nil
}

func copyFile(src, dst string) error {
    source, err := os.Open(src)
    if err != nil {
        return err
    }
    defer source.Close()

    destination, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer destination.Close()

    _, err = io.Copy(destination, source)
    return err
}