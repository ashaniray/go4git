# TOOLS
- Each tool will have a main(). Build with ```go build <your_tool>.go utils.go```. This will produce ```<your_tool>``` binary
- Update ```Makefile``` for the new tool:
  - new target for builing of new tool. Refer to target ```unzlib``` in ```Makefile```
  - update ```all``` target
  - update ```install``` target
- The ```main``` will ofen take an argument filename to read from. If not specified it will read from stdin. Like **sort**, **cat**, etc. Call ```getArgInputFile()``` method in utils.go to obtail the File* for this purpose
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

### gensha1
Generates hash of **blob**

#### Usage

```
$ gensha1 <filename> | echo $(hexdump -ve '1/1 "%.2x"')
```

Compare with
```
$ echo 'test content' | git hash-object --stdin

$ echo 'test content' | gensha1 | echo $(hexdump -ve '1/1 "%.2x"')

$ echo -e 'blob 14\0Hello, World!' | shasum

```

---
### obj2file
Converts a loose object (hash) to file path.

#### Usage
```
$ cat `obj2file <hash>`  |unzlib
```


## TODO
- New utility **ls-tree** to list the contents of a tree

- New utility **tree2fs** to convert a tree into a folder structure in the file-system.
