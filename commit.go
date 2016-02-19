package go4git

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
)

type CommitFields map[string]string

type Commit struct {
	id        string
	tree      string
	parent    string
	author    *Person
	committer *Person
	message   string
}

func (c *Commit) Id() string {
	return c.id
}

func (c *Commit) Tree() string {
	return c.tree
}

func (c *Commit) Parent() string {
	return c.parent
}

func (c *Commit) Author() *Person {
	return c.author
}

func (c *Commit) Committer() *Person {
	return c.committer
}

func (c *Commit) Message() string {
	return c.message
}

func (c *Commit) HasParent() bool {
	return len(c.parent) != 0
}

func (c *Commit) String() string {
	buff := new(bytes.Buffer)
	fmt.Fprintf(buff, "Tree:      %s\n", c.tree)
	fmt.Fprintf(buff, "Parent:    %s\n", c.parent)
	fmt.Fprintf(buff, "Author:    %s\n", c.author)
	fmt.Fprintf(buff, "Committer: %s\n", c.committer)
	fmt.Fprintf(buff, "Message:   %s\n", c.message)

	return string(buff.Bytes())
}

func kvPair(s string) (string, string) {
	xs := strings.Split(s, " ")
	return xs[0], strings.Join(xs[1:], " ")
}

func parseCommitBody(buff *bytes.Buffer, size int) (CommitFields, error) {
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

	commit.tree = cf["tree"]
	commit.author = parsePerson(cf["author"])
	commit.committer = parsePerson(cf["committer"])
	commit.message = cf["message"]

	parent, ok := cf["parent"]

	if ok {
		commit.parent = parent
	}
	return commit
}

func ParseCommit(in io.Reader) (*Commit, error) {

	buff := new(bytes.Buffer)
	buff.ReadFrom(in)

	size, typ, err := parseHeader(buff)

	if err != nil {
		return nil, err
	}

	if typ != "commit" {
		return nil, errors.New("not a commit")
	}

	body, err := parseCommitBody(buff, size)

	if err != nil {
		return nil, err
	}

	return body.ToCommit(), nil
}
