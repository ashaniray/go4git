package main

import (
	"path"
	"path/filepath"
)

type Repository struct {
	isBare bool
	gitDir string
}

func NewRepository(path string) (*Repository, error) {
	bare := isBareRepo(path)
	gd, err := getGitDir(path)

	if err != nil {
		return nil, err
	}
	return &Repository{isBare: bare, gitDir: gd}, nil
}

func (r *Repository) IsBare() bool {
	return r.isBare
}

func (r *Repository) WorkDir() string {
	if r.isBare {
		return ""
	} else {
		return path.Dir(r.gitDir)
	}
}

func (r *Repository) Path() string {
	return r.gitDir
}

func (r *Repository) LooseObjPath(sha string) string {
	return filepath.Join(r.gitDir, "objects", sha[0:2], sha[2:])
}


