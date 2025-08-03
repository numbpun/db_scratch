

The function `SaveFile` in ../Examples/save.go uses filesystem as KV
FileName = "key"
File Data = "value"

Structure is {path: File (Data)}

`fsync` syscall that requests/confirms that file data is written

`O_TRUNC`Flag: Discards the previous data if file exists

`O_CREATE` Flag: Creates file if doesn't exist

Required cuz there may be many levels of buffering:
- OS page cache
- on-device RAM

Successful `fsync` = permanent data on final medium

`Sync()` calls `fsync`

Method fails atomicity and durability requirements:

- Truncating a file means possible to end up with empty file
- Not truncating means possible to end up with a half written file

And this is the problem of **Updating Data in-place**

This is why we need a temp file or buffer which temprorarily stores the the data or stores the whole data and switches to it.

Switch operation is `rename()` in filesystem specifically designed to make atomic files possible.
Therefore, we have function `SaveFile2`

### Types of Atomicity

1. **Power-Loss-Atomicity**: Will a reader observe a bad state after a crash

2. **Readers-writers atomic**: Readers see bad state with concurrent writer

`SaveFile1` failed both atomicity; reader would have been able to observe the empty file with/without crash

`SaveFile2` is readers-writes atomics only (NOT power loss atomic. Why?)

#### Why does renaming work?
- Filesystem use maps {File Name: File Data}, so replacing or "renaming" (switching) simply changes the pointer (file name) of old data to new data (therefore old data is untouched).

- Mapping is just "Directory" and is many-to-one i.e. multiple names can reference the same file even from different directories (*hard link*). 

- File with 0 references is automatically deleted.

- Atomicity and durability of `rename()` depends on directory updates but updation of directory is only *readers-writers atomic* and not power-loss atomic or durable. Hence `SaveFile2` is still missing some component

### `fsync` issue
Creation/ renaming a file updates containing directory.

Therefore, to make directories durable, `fsync` can be called on them to using directory's handle / *file descriptor*.

However, if `fsync` fails -> DB update fails -> What if read file afterwards?

- We might get new data even if fsync failed due to OS Page cache (What is it?). This behavior is filesystem dependent (Explain fsync failure scenario)

# Atomic log updates with checksums

Checksums help with detecting silent data corruption. It basically checks the size of data and help detect the corrupted last entry and ignore it when recovering from a crash