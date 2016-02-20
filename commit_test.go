package go4git

import (
	"bytes"
	"testing"
)

func TestParseCommit(t *testing.T) {
	rd := bytes.NewReader(COMMIT_DATA)

	commit, err := ParseCommit(rd)

	if err != nil {
		t.Fatalf("Failed to parse commit - [%s]", err)
	}

	if got, want := commit.Tree(), "bb8973209e6a81bca6ae065db330b04c225b1d45"; got != want {
		t.Errorf("Tree = %s; want %s", got, want)
	}

	if got, want := commit.Message(), "first commit"; got != want {
		t.Errorf("Message = %s; want %s", got, want)
	}

	if got, want := commit.Author().Name, "Mrs. Eldridge McDermott"; got != want {
		t.Errorf("Author name = %s; want %s", got, want)
	}

	if got, want := commit.Author().Email, "sandy_schmidt@stiedemann.biz"; got != want {
		t.Errorf("Author email = %s; want %s", got, want)
	}

	if got, want := commit.Committer().Name, "Mrs. Eldridge McDermott"; got != want {
		t.Errorf("Committer name = %s; want %s", got, want)
	}

	if got, want := commit.Committer().Email, "sandy_schmidt@stiedemann.biz"; got != want {
		t.Errorf("Committer email = %s; want %s", got, want)
	}

	if commit.HasParent() {
		t.Errorf("This commit should not have a parent")
	}

}

func TestParseCommitWithParent(t *testing.T) {
	rd := bytes.NewReader(COMMIT_DATA_WP)

	commit, err := ParseCommit(rd)

	if err != nil {
		t.Fatalf("Failed to parse commit - [%s]", err)
	}

	if got, want := commit.Tree(), "c2c516707c7aa0bd3a6d64284c2f04becf76422d"; got != want {
		t.Errorf("Tree = %s; want %s", got, want)
	}

	if got, want := commit.Message(), "If we generate the port, we can get to the XSS firewall through the neural PCI feed!"; got != want {
		t.Errorf("Message = %s; want %s", got, want)
	}

	if got, want := commit.Author().Name, "Bettye Leffler"; got != want {
		t.Errorf("Author name = %s; want %s", got, want)
	}

	if got, want := commit.Author().Email, "carol@weinatromaguera.biz"; got != want {
		t.Errorf("Author email = %s; want %s", got, want)
	}

	if got, want := commit.Committer().Name, "Bettye Leffler"; got != want {
		t.Errorf("Committer name = %s; want %s", got, want)
	}

	if got, want := commit.Committer().Email, "carol@weinatromaguera.biz"; got != want {
		t.Errorf("Committer email = %s; want %s", got, want)
	}

	if got, want := commit.Parent(), "5e493b5a14202fa5cb40f879ecc0de9b8a8f453a"; got != want {
		t.Errorf("Parent = %s; want %s", got, want)
	}

}
