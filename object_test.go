package go4git

import (
	"bytes"
	"testing"
)

func TestCommitHdr(t *testing.T) {
	buff := bytes.NewBuffer(COMMIT_DATA)
	size, typ, err := parseHeader(buff)

	if err != nil {
		t.Fatalf("Falied to parse commit header - [%s]", err)
	}

	if got, want := size, 221; got != want {
			t.Errorf("size = %d; want %d", got, want)
	}

	if got, want := typ, "commit"; got != want {
			t.Errorf("type = %s; want %s", got, want)
	}
}

func TestTagHdr(t *testing.T) {
	buff := bytes.NewBuffer(TAG_DATA)
	size, typ, err := parseHeader(buff)

	if err != nil {
		t.Fatalf("Failed to parse tag header - [%s]", err)
	}

	if got, want := size, 224; got != want {
			t.Errorf("size = %d; want %d", got, want)
	}

	if got, want := typ, "tag"; got != want {
			t.Errorf("type = %s; want %s", got, want)
	}
}

func TestTreeHdr(t *testing.T) {
	buff := bytes.NewBuffer(TREE_DATA)
	size, typ, err := parseHeader(buff)

	if err != nil {
		t.Fatalf("Failed to parse tree header - [%s]", err)
	}

	if got, want := size, 127; got != want {
			t.Errorf("size = %d; want %d", got, want)
	}

	if got, want := typ, "tree"; got != want {
			t.Errorf("type = %s; want %s", got, want)
	}
}

func TestBlobHdr(t *testing.T) {
	buff := bytes.NewBuffer(BLOB_DATA)
	size, typ, err := parseHeader(buff)

	if err != nil {
		t.Fatalf("Failed to parse blob header - [%s]", err)
	}

	if got, want := size, 105; got != want {
			t.Errorf("size = %d; want %d", got, want)
	}

	if got, want := typ, "blob"; got != want {
			t.Errorf("type = %s; want %s", got, want)
	}
}
