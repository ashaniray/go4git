package main

import(
	"fmt"
	"io"
	"bytes"
	"strings"
	"errors"
	"strconv"
)

type CommitFields map[string]string

type Commit struct {
	Tree string
	Parent string
	Author string
	Committer string
	Message   string
}

func (c *Commit) String() string {
	buff := new(bytes.Buffer)
	fmt.Fprintf(buff, "Tree:      %s\n", c.Tree)
	fmt.Fprintf(buff, "Parent:    %s\n", c.Parent)
	fmt.Fprintf(buff, "Author:    %s\n", c.Author)
	fmt.Fprintf(buff, "Committer: %s\n", c.Committer)
	fmt.Fprintf(buff, "Message:   %s\n", c.Message)

	return string(buff.Bytes())
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

func kvPair(s string) (string, string) {
	xs := strings.Split(s, " ")
	return xs[0], strings.Join(xs[1:], " ")
}

func parseBody(buff *bytes.Buffer, size int) (CommitFields, error) {
	body := string(buff.Next(size))

	lines := strings.Split(body, "\n")

	fields := make(map[string]string)

	// TODO: add validations and return error
	for i, l := range lines {
		if len(l) == 0 {
			fields["message"] = strings.Trim(strings.Join(lines[i:], "\n"), "\n")
			break
		} else {
			k, v := kvPair(l)
			fields[k] = v
		}
	}

	return fields, nil
}

func (cf CommitFields) ToCommit() *Commit {
	commit := new(Commit)

	commit.Tree      = cf["tree"]
	commit.Author    = cf["author"]
	commit.Committer = cf["committer"]
	commit.Message   = cf["message"]

	parent, ok := cf["parent"] 

	if ok {
		commit.Parent = parent
	}
	return commit;
}

func ParseCommit(in io.Reader) (*Commit, error) {

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

	return body.ToCommit(), nil
}

