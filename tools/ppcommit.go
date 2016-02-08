package main

import(
	"fmt"
	"os"
	"flag"
	"io"
	"bytes"
	"strings"
	"errors"
	"strconv"
)

var repoRoot = flag.String("d", ".", "path to repo")


type Commit struct {
	Tree string
	Parent string
	Author string
	Committer string
}

func parseHeader(buff *bytes.Buffer) (int, error) {
	header, err := buff.ReadString(0)

	if err != nil {
		return 0, err
	}

	xs := strings.Split(header, " ")
	objType, objSize := xs[0], xs[1]

	if objType != "commit" {
		return 0, errors.New("not a commit object")
	}

	objSize = objSize[:len(objSize)-1] // remove trailing null

	size, err := strconv.Atoi(objSize)

	if err != nil {
		return 0, err
	}

	return size, nil
}

func parseBody(buff *bytes.Buffer, size int) (*Commit, error) {
	body := string(buff.Next(size))

	lines := strings.Split(body, "\n")

	for _, l := range lines {
		fmt.Println("LINE:", l)
	}
	

	return nil, nil
}

func ParseCommit(in io.Reader) (*Commit, error) {

	commit := new(Commit)

	buff := new(bytes.Buffer)
	buff.ReadFrom(in)

	size, err := parseHeader(buff) 

	if err != nil {
		return nil, err
	}

	body, err := parseBody(buff, size)

	if err != nil {
		return nil, err
	}

	fmt.Println("BODY:", body)

	return commit, nil
}

func main() {
	flag.Parse()

	f, err := GetArgInputFile()

	cmt, err := ParseCommit(f)

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)	
		return
	}

	fmt.Fprintln(os.Stdout, cmt)
}
