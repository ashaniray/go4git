package go4git

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
)

type Tag struct {
	name       string
	targetId   string
	targetType string
	tagger     *Person
	message    string
}

func (t *Tag) String() string {
	buff := new(bytes.Buffer)

	fmt.Fprintf(buff, "Name:             %s\n", t.name)
	fmt.Fprintf(buff, "Target ID:        %s\n", t.targetId)
	fmt.Fprintf(buff, "Target Type:      %s\n", t.targetType)
	fmt.Fprintf(buff, "Tagger:           %s\n", t.tagger)
	fmt.Fprintf(buff, "Message:          %s\n", t.message)

	return string(buff.Bytes())
}

func (t Tag) Name() string {
	return t.name
}

func (t Tag) TargetId() string {
	return t.targetId
}

func (t Tag) TargetType() string {
	return t.targetType
}

func (t Tag) Tagger() *Person {
	return t.tagger
}

func (t Tag) Message() string {
	return t.message
}

type TagFields map[string]string

func (tf TagFields) ToTag() *Tag {
	tag := new(Tag)
	tag.name = tf["tag"]
	tag.targetId = tf["object"]
	tag.targetType = tf["type"]
	tag.tagger = parsePerson(tf["tagger"])
	tag.message = tf["message"]
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

func ParseTag(in io.Reader) (*Tag, error) {

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
