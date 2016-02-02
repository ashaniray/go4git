
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
```