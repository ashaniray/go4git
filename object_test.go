package go4git

import (
	"bytes"
	"testing"
)

func TestCommitHdr(t *testing.T) {
	buff := bytes.NewBuffer(COMMIT_DATA)
	size, typ, err := parseHeader(buff)

	exSize := 221
	exTyp := "commit"

	if err != nil {
		t.Errorf("Expected no error while parsing commit header but got %s", err)
	}

	if size != exSize {
		t.Errorf("Expected size to be %d but got %d", exSize, size)
	}

	if typ != exTyp {
		t.Errorf("Expected type to be %s but got %s", exTyp, typ)
	}
}

func TestTagHdr(t *testing.T) {
	buff := bytes.NewBuffer(TAG_DATA)
	size, typ, err := parseHeader(buff)

	exSize := 224
	exTyp := "tag"

	if err != nil {
		t.Errorf("Expected no error while parsing tag header but got %s", err)
	}

	if size != exSize {
		t.Errorf("Expected size to be %d but got %d", exSize, size)
	}

	if typ != exTyp {
		t.Errorf("Expected type to be %s but got %s", exTyp, typ)
	}
}

func TestTreeHdr(t *testing.T) {
	buff := bytes.NewBuffer(TREE_DATA)
	size, typ, err := parseHeader(buff)

	exSize := 127
	exTyp := "tree"

	if err != nil {
		t.Errorf("Expected no error while parsing tree header but got %s", err)
	}

	if size != exSize {
		t.Errorf("Expected size to be %d but got %d", exSize, size)
	}

	if typ != exTyp {
		t.Errorf("Expected type to be %s but got %s", exTyp, typ)
	}
}

func TestBlobHdr(t *testing.T) {
	buff := bytes.NewBuffer(BLOB_DATA)
	size, typ, err := parseHeader(buff)

	exSize := 105
	exTyp := "blob"

	if err != nil {
		t.Errorf("Expected no error while parsing blob header but got %s", err)
	}

	if size != exSize {
		t.Errorf("Expected size to be %d but got %d", exSize, size)
	}

	if typ != exTyp {
		t.Errorf("Expected type to be %s but got %s", exTyp, typ)
	}
}
