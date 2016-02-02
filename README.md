# go4git
Git library in pure GO (Under construction)

# Usage

#### Open an existing repository.
```
repo := go4git.NewRepository("path/to/my/repository")
```


#### Create a new bare repository

```
bareRepo := go4git.NewBareRepository(".")
```

#### Accessing a Repository

```
// Returns `true` if the given SHA1 exist in this repository
repo.Exists("07b44cbda23b726e5d54e2ef383495922c024202")
```

#### Query repository state
```
repo.IsBare()
repo.IsEmpty()
repo.IsHeadUnborn()
repo.IsHeadDetached()
```


#### Path accessors
```
repo.Path()
repo.Workdir()
```

#### The HEAD of the repository.
```
ref := repo.Head()
```

#### Properties of ref
```
ref.Name()
ref.Target()
```

#### Reading an object
```
obj := repo.Read("a0ae5566e3c8a3bddffab21022056f0b5e03ef07")
obj.Length()
obj.Data()
obj.Type()
```



#### Writing to a Repository
```
sha := repo.Write([]byte("some content."), go4git.Blob)
```




#### Commit Objects
```
commit := repo.Lookup('a0ae5566e3c8a3bddffab21022056f0b5e03ef07')

commit.Message()
commit.Time()
commit.Author()
commit.Tree()
commit.Parents()
```

#### Tag Objects

```
tag, err:= repo.LookupTag('a0ae5566e3c8a3bddffab21022056f0b5e03ef07')

tag.Target()
tag.Target().Oid
tag.Target().Type // can be one of go4git.Commit, go4git.Tag, go4git.Blob
tag.Name()
tag.Message()
tag.Tagger()
```


#### Tree Objects
```
tree, err:= repo.LookupTree('779fbb1e17e666832773a9825875300ea736c2da')

tree.Count()
tree[0]
```


#### Blob Objects

```
blob := repo.LookupBlob('e1253910439ea902cf49be8a9f02f3c08d89ac73')
blob.Content() // Gives the content of the blob.
```


#### Manipulating git index
```
index := go4git.NewIndex(path)


index.Reload()          // Reload the index file from disk.
count = index.Count()   // Get the count of index entries.
index.Entries()         // Get the collection of index entries.

// Iterating over index entries.
for i := range index {
  fmt.Println(i)
}

index.Entry(path)         // Get a particular entry in the index.
index.Remove(path)        // Remove from staging
index.Add(entry)          // Add to staging. Also updates existing entry if there is one.
index.AddFromPath(path)   // Add to staging. Creates entry from file in path, updates the index.

```


#### References
```
ref := repo.Reference("refs/heads/master")

ref.Target().Id // SHA1 hash
ref.Type()      // go4git.Direct
ref.Name()      // "refs/heads/master"

// Iterate over all references:

refs := repo.References()

for ref := range refs {
  fmt.Println(ref)
}



// Iterate only over references that match the given pattern (glob):

refs := repo.ReferencesByGlob("refs/tags/*")

for ref := range refs {
  fmt.Println(ref)
}
```

#### Create, update, rename or delete a reference

```
ref := repo.CreateReference("refs/heads/unit_test", some_commit_sha)

repo.UpdateReference(ref, newSha)
repo.UpdateReferenceByName("refs/heads/unit_test", newSha)

repo.RenameReference(ref, "refs/heads/blead") 
repo.RenameReferenceByName("refs/heads/unit_test", "refs/heads/blead")

repo.DeleteReference(ref)
repo.DeleteReferenceByName("refs/heads/unit_test")
```

#### Access the reflog for any branch:

```
ref      := repo.Reference("refs/heads/master")
reflog   := ref.Log()
entry    := reflog[0]      // Get the first entry

entry.OldId
entry.NewId
entry.Message
entry.Committer
```

### Branches

#### Iterate over all branches:

```
branches := repo.LocalBranches()

for branch := range branches {       // ["master"]
  fmt.Println(branch.Name())
}

branches := repo.RemoteBranches()

for branch := range branches {       // ["origin/HEAD", "origin/master", "origin/packed"]
  fmt.Println(branch.Name())
}
```
#### Look up branches and get attributes

```
branch = repo.Branch("master")
branch.Name()           // "master"
branch.CanonicalName()  // "refs/heads/master"
```
#### Look up the id for the target of a branch:

branch.Target().Id // "36060c58702ed4c2a40832c51758d5344201d89a"

#### Create and delete branch

```
branch = repo.CreateBranch("test_branch", "HEAD")

repo.branches.RenameBranchByName("test_branch", "new_branch")
repo.RenameBranchByCName("refs/heads/test_branch", "new_branch")
repo.RenameBranch(ref, "new_branch")

repo.DeleteBranchByName("test_branch")
repo.DeleteBranchByCName("refs/heads/test_branch")
repo.DeleteBranch(ref)
```
