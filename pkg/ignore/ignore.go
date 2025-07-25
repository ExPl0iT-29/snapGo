package ignore

import (
    "os"
    gitignore "github.com/crackcomm/go-gitignore"
)

func LoadIgnorePatterns(ignoreFile string) (*gitignore.GitIgnore, error) {
    if _, err := os.Stat(ignoreFile); os.IsNotExist(err) {
        return gitignore.CompileIgnoreLines(), nil
    }
    return gitignore.CompileIgnoreFile(ignoreFile)
}

func ShouldIgnore(path string, ignorer *gitignore.GitIgnore) bool {
    return ignorer.MatchesPath(path)
}