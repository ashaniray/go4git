
## Using the tools
### unzlib
```
cat <git_object_file_path> | unzlib
```
### git tools
```
$ git ls-tree <git_object>
$ git show -s --pretty=raw <git_object>
$ git cat-file tag v1.5.0
$ git cat-file -p <git_object>
$ git cat-file tree <git_object>
$ git update-index --add  --cache-info 100644 <hash> <filename>
$ git write-tree
$ git commit-tree <hash> -m "Message"
$ echo <commit_hash> > .git/refs/heads/master

```

### Minimal git repository
```
.git/
├── HEAD
├── objects
└── refs
    └── heads
```
