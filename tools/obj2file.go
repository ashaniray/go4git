package main

import (
	"os"
	"fmt"
	"flag"
	"path/filepath"
	"errors"
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

func GitDir(root string) (string, error) {
	switch {
	case IsRepo(*repoRoot):
		return filepath.Join(root, ".git"), nil
	case IsBareRepo(*repoRoot):
		return root, nil
	default:
		return "", errors.New("not a git repository")
	}
}

func GetObjPath(sha string, root string) (string, error) {
	gd, err := GitDir(root)

	if err != nil {
		return "", err
	}

	return filepath.Join(gd,"objects", sha[0:2], sha[2:]), nil
}


var repoRoot = flag.String("d", ".", "path to repository root")

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "ERROR: Invalid arguments")
		return
	}

	p, err := GetObjPath(flag.Arg(0), *repoRoot)

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
	} else {
		fmt.Fprintf(os.Stdout, p)
	}

}
