package main

import (
	"fmt"
	"os"
	"flag"
	"path/filepath"
	"io/ioutil"
	"strings"
)

var repoRoot = flag.String("d", ".", "path to a git repo")

func getHead(gitDir string) (string, error) {
	head := filepath.Join(gitDir, "HEAD")
	data, err := ioutil.ReadFile(head)

	if err != nil {
		return "", err
	}

	line := strings.Trim(string(data), "\r\n")
	cmpts := strings.Split(line, ":")
	ref := strings.Trim(cmpts[1], " \r\n")

	refPath := filepath.Join(gitDir, ref)

	refSha, err := ioutil.ReadFile(refPath)
	
	if err != nil {
		return "", err
	}
	

	return strings.Trim(string(refSha), " \r\n"), nil
}

func main() {
	flag.Parse()

	repof, err := GitDir(*repoRoot)

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		return
	}

	headSha, err := getHead(repof)

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		return
	}

	fmt.Fprintf(os.Stdout, headSha)
}
