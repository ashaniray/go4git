package go4git

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func refContent(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return strings.Trim(string(b), "\r\n"), nil
}

func (r *Repository) References() ([]Reference, error) {
	refs := make([]Reference, 0)

	headsDir := filepath.Join(r.gitDir, "refs", "heads")
	tagsDir := filepath.Join(r.gitDir, "refs", "tags")

	heads, err := ioutil.ReadDir(headsDir)

	if err != nil {
		return refs, err
	}

	for _, h := range heads {
		n := filepath.Join("refs", "heads", h.Name())
		rpath := filepath.Join(r.gitDir, n)

		targetSha, err := refContent(rpath)
		if err != nil {
			return refs, err
		}
		ref := Reference{name: n, targetId: targetSha}
		refs = append(refs, ref)
	}

	tags, err := ioutil.ReadDir(tagsDir)

	if err != nil {
		return refs, err
	}

	for _, t := range tags {
		n := filepath.Join("refs", "tags", t.Name())
		rpath := filepath.Join(r.gitDir, n)
		targetSha, err := refContent(rpath)

		if err != nil {
			return refs, err
		}
		ref := Reference{name: n, targetId: targetSha}
		refs = append(refs, ref)
	}
	return refs, nil
}

type Reference struct {
	name     string
	targetId string
}

func (r Reference) IsBranch() bool {
	return false
}

func (r Reference) IsTag() bool {
	return false
}

func (r Reference) IsRemote() bool {
	return false
}

func (r Reference) HasLog() bool {
	return false
}

func (r Reference) String() string {
	return fmt.Sprintf("%s %s", r.targetId, r.name)
}

func (r Reference) Type() string {
	return ""
}

func (r Reference) Name() string {
	return ""
}
