package go4git

import (
	"bytes"
	"testing"
)

func TestParseTag(t *testing.T) {
	rd := bytes.NewReader(TAG_DATA)

	tag, err := ParseTag(rd)

	if err != nil {
		t.Fatalf("Failed to parse tag - [%s]", err)
	}

	if got, want := tag.Name(), "v02.12"; got != want {
		t.Errorf("name = %s; want %s", got, want)
	}

	if got, want := tag.TargetId(), "0f9e744ee1dbbedeac469ee0625db1fefece53db"; got != want {
		t.Errorf("target id = %s; want %s", got, want)
	}

	if got, want := tag.TargetType(), "commit"; got != want {
		t.Errorf("target type = %s; want %s", got, want)
	}

	if got, want := tag.Message(), "You can't connect the driver without quantifying the solid state AGP bandwidth!"; got != want {
		t.Errorf("message = %s; want %s", got, want)
	}
}
