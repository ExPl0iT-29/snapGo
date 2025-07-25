package ignore

import (
    "github.com/sabhiram/go-gitignore"
    "os"
)

func LoadIgnorePatterns(ignoreFile string) (*gitignore.GitIgnore, error) {
    if _, err := os.Stat(ignoreFile); os.IsNotExist(err) {
        return gitignore.CompileIgnoreLines(), nil
    }
    return gitignore.ParseIgnoreFile(ignoreFile)
}

func ShouldIgnore(path string, ignorer *gitignore.GitIgnore) bool {
    return ignorer.MatchesPath(path)
}