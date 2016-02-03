# TOOLS
- Each tool will have a main(). Build with ```go build <your_tool>.go util.go```. This will produce ```<your_tool>``` binary 
- The main will ofen take an argument filename to read from. If not specified it will read from stdin. Like **sort**, **cat**, etc. Call ```getArgInputFile()``` method in utils.go to obtail the File* for this purpose
- Place your methods in ```utils.go``` and call the method from main.
- The main will only prcoess the command line and call another function with proper arguments - so that the function 
can also be called from other functions. See ```unzlib.go``` as an example
- When adding/modifying code in utils.go demarcate the section you are modifying/adding with a marker to avoid merge conflicts. All your changes should be inside the marked section. E.g.
```
..some code
/////Changes by <your_name> ////
<Changes by your_name goes here>
func foo() {
}
func xyz() {
}
/////End of changes by <your_name>

//////Changes by <some_other_name> ..
<Changes by some other name goes here>
func newFoo() {
}
//// End of changes by <some_other_name> //
```

### unzlib
Decompressed a zlib compressed data.
```
cat <filename> | unzlib
```
OR
```
unzlib <filename>
```

## TODO
- New utility **gen-sha1** to generate hash of content (optional filename as arg, if no arg then read from stdin)
- New utility **ls-tree** to list the contents of a tree
- New utility **expand-tree** to expand tree util
- New utility **obj2file** to convert object (hash) has to file path. Example usage: cat `obj2file <hash>` | unzlib
- New utility **init-tree** to convert a tree into a .git structure with hashes etc. Like git init.
