# TOOLS
- Each tool will have a main(). 
- The main will generally take an argument (e.g. filename). If not specified it will read from stdin. Like **sort**, **cat**, etc. It might be possible to have a common function for this.
- The main will only prcoess the command line and call another function with proper arguments - so that the function 
can also be called from other functions.

### unzlib
Decompressed a zlib compressed data.
```
cat <filename> | unzlib
```

## TODO
- Add file name as arg to **unzlib** (if no arg then read from stdin). Currently it always reads from stdin
- New utility **gen-sha1** to generate hash of content (optional filename as arg, if no arg then read from stdin)
- New utility **ls-tree** to list the contents of a tree
- New utility **expand-tree** to expand tree util
- New utility **obj2file** to convert object (hash) has to file path. Example usage: cat `obj2file <hash>` | unzlib
- New utility **init-tree** to convert a tree into a .git structure with hashes etc. Like git init.
