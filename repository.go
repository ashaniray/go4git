package go4git

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"bytes"
	"bufio"
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

func (r *Repository) LooseObjects() ([]string, error) {

	objects := make([]string, 0)

	objDir := filepath.Join(r.gitDir, "objects")

	entries, err := ioutil.ReadDir(objDir)

	if err != nil {
		return objects, err
	}

	for _, e := range entries {
		if e.IsDir() && len(e.Name()) == 2 {
			subd := filepath.Join(objDir, e.Name())
			subEntries, err := ioutil.ReadDir(subd)
			if err != nil {
				return objects, err
			}

			for _, se := range subEntries {
				obj := e.Name() + se.Name()
				objects = append(objects, obj)
			}
		}
	}

	return objects, nil

}

func (r *Repository) LookupCommit(sha string) (*Commit, error) {
	lobjp := r.LooseObjPath(sha)
	f, err := os.Open(lobjp)

	if err != nil {
		return nil, err
	}

	buff := new(bytes.Buffer)
	err = Unzlib(bufio.NewReader(f), buff)

	if err != nil {
		return nil, err
	}
	commit, err := ParseCommit(buff)

	if err != nil {
		return nil, err
	}

	commit.Id = sha
	return commit, nil
}


func InitAt(root string, bare bool) error {

	var gitDir string

	if bare {
		gitDir = root
	} else {
		gitDir = filepath.Join(root, ".git")
	}

	err := createFolders(gitDir)

	if err != nil {
		return err
	}

	head := filepath.Join(gitDir, "HEAD")
	err = ioutil.WriteFile(head, []byte("ref: refs/heads/master"), os.ModePerm)

	return err
}
