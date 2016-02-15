package main

import (
	"io"
	"bytes"
	"strings"
	"errors"
	"fmt"
)


type Tag struct {
	Name       string
	TargetId   string
	TargetType string
	Tagger     *Person	
	Message    string
}

func (t *Tag) String() string {
	buff := new(bytes.Buffer)

	fmt.Fprintf(buff, "Name:             %s\n", t.Name)
	fmt.Fprintf(buff, "Target ID:        %s\n", t.TargetId)
	fmt.Fprintf(buff, "Target Type:      %s\n", t.TargetType)
	fmt.Fprintf(buff, "Tagger:           %s\n", t.Tagger)
	fmt.Fprintf(buff, "Message:          %s\n", t.Message)

	return string(buff.Bytes())
}

type TagFields map[string]string

func (tf TagFields) ToTag() *Tag {
	tag := new(Tag)
	tag.Name         = tf["tag"]
	tag.TargetId     = tf["object"]
	tag.TargetType   = tf["type"]
	tag.Tagger       = parsePerson(tf["tagger"])
	tag.Message      = tf["message"]
	return tag
}

func parseTagBody(buff *bytes.Buffer, size int) (TagFields, error) {
	body := string(buff.Next(size))

	lines := strings.Split(body, "\n")

	fields := make(map[string]string)

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

func parseTag(in io.Reader) (*Tag, error) {

	buff := new(bytes.Buffer)
	buff.ReadFrom(in)

	size, typ, err := parseHeader(buff)

	if err != nil {
		return nil, err
	}

	if typ != "tag" {
		return nil, errors.New("not a tag")
	}

	body, err := parseTagBody(buff, size) 

	if err != nil {
		return nil, err
	}

	return body.ToTag(), nil

}
