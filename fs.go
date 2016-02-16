package go4git

import (
	"errors"
	"os"
	"path/filepath"
)

func fileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else {
		return false
	}
}

func folderExists(path string) bool {
	if s, err := os.Stat(path); err == nil {
		return s.IsDir()
	} else {
		return false
	}
}

func isRepo(root string) bool {
	dotGit := filepath.Join(root, ".git")
	return folderExists(dotGit)
}

func isBareRepo(root string) bool {
	objFolder := filepath.Join(root, "objects")
	headFile := filepath.Join(root, "HEAD")

	return folderExists(objFolder) && fileExists(headFile)
}

func getGitDir(root string) (string, error) {
	absRoot, err := filepath.Abs(root)

	if err != nil {
		return "", err
	}

	switch {
	case isRepo(absRoot):
		return filepath.Join(absRoot, ".git"), nil
	case isBareRepo(absRoot):
		return absRoot, nil
	default:
		return "", errors.New("not a git repository")
	}
}

func createFolders(root string) error {
	var folders = []string{
		"branches", "hooks", "info",
		"objects", "objects/info", "objects/pack",
		"refs", "refs/heads", "refs/tags",
	}

	for _, f := range folders {
		folderPath := filepath.Join(root, f)
		err := os.MkdirAll(folderPath, os.ModePerm)

		if err != nil {
			return err
		}
	}

	return nil
}
