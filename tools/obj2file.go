package main

import (
	"os"
	"fmt"
	"flag"
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

func IsRepo(root string) bool {
	dotGit := filepath.Join(root, ".git")
	return folderExists(dotGit)
}

func IsBareRepo(root string) bool {
	objFolder := filepath.Join(root, "objects")
	headFile := filepath.Join(root, "HEAD")

	return folderExists(objFolder) && fileExists(headFile)
}


var repoRoot = flag.String("d", ".", "path to repository root")

func main() {
	flag.Parse()

	// dir, _ := os.Getwd()

	switch {
	case IsRepo(*repoRoot):
		fmt.Println("repo")
	case IsBareRepo(*repoRoot):
		fmt.Println("bare repo")
	}
	// is it a normal repo
	// os.IsExists(dir + 

	// or a bare repo 

	// or not a repo at all

	
}
