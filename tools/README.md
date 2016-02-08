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
---

### unzlib
Decompressed a zlib compressed data.
```
cat <filename> | unzlib
```
OR
```
unzlib <filename>
```
---
### zlib
zlib compress data

#### Usage

```
$ cat <filename> | zlib
```
OR
```
$ zlib <filename>
```
---
### gensha1
Generates hash of (**blob** by default)

#### Usage

```
$ gensha1 <filename> | echo $(hexdump -ve '1/1 "%.2x"')
$ echo "<contents_of_commit>" | gensha1 -t commit | echo $(hexdump -ve '1/1 "%.2x"')
```

Compare with
```
$ echo 'test content' | git hash-object --stdin

$ echo 'test content' | gensha1 | echo $(hexdump -ve '1/1 "%.2x"')

$ echo -e 'blob 14\0Hello, World!' | shasum

```
---

### lstree
Lists the details of a tree object

#### Usage
```
$ cat <filename> | unzlib | lstree
```
Compare the following:
```
$ git cat-file -p <hash_of_a_tree>

$ cat `obj2file <hash_of_a_tree>` |  unzlib | lstree
```
---
### lstype
Lists the type object

#### Usage
```
$ cat <filename> | unzlib | lstype
```
Compare the following:
```
$ git cat-file -t <hash_of_object>

$ $ cat `obj2file <hash_of_object>` |  unzlib | lstype
```

---
### obj2file
Converts a loose object (hash) to file path.

#### Usage
```
$ cat `obj2file <hash>`  |unzlib
```
---
### lslobj
Lists all loose objects in a repository.

#### Usage

```
$ lslobj -d /path/to/repo
```
---


### initr
Creates an empty [bare] git repository

#### Usage

```
$ initr /path/to/repo
$ initr -bare /path/to/repo
```
---


### ppcommit
Pretty print a commit object, reading the decompressed commit object from stdin or file.

#### Usage

```
$ cat `obj2file <commit-hash>` | unzlib | ppcommit
```


## TODO
- **tree2fs** to convert a tree into a folder structure in the file-system.
- **lscommits** list all commits of the current branch
- **fs2tree** convert a folder structure to a tree object
- **unpack** Unpack objects from a packed archive
- **pack** Create a packed archive of objects
