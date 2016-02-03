# TOOLS
### unzlib
Decompressed a zlib compressed data.
```
cat <filename> | unzlib
```

## TODO
- Add file name as arg to **unzlib** (if no arg then read from stdin). Currently it always reads from stdin
- New utility **gen-sha1** to generate hash of content (optional filename as arg, if no arg then read from stdin)
- New utility **ls-tree** to list the contents of a tree
- New utility **create-tree** to expand tree util
- New utility **obj2file** to convert object (hash) has to file path. Example usage: cat `obj2file <hash>` | unzlib
