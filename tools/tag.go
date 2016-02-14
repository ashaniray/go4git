package main

import (
	"io"
	"bytes"
)

type Tag struct {
	Name       string
	TargetId   string
	TargetType int
	Tagger     *Person	
}

func parseTag(in io.Reader) (*Tag, error) {

	buff := new(bytes.Buffer)
	buff.ReadFrom(in)

	parseHeader(buff)

	return nil, nil
}
