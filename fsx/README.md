A simple and intuitive Go library that provides convenient functions for file system operations. It offers a clean API for common tasks like reading files, manipulating paths, and watching for system events.

```go
config, err := fsx.ReadFileJsonAs[Config]("./config.json")
if err != nil {
  panic(err)
}

fsx.WatchRecursive(context.Background(), config.BaseDir, func (e fsx.Event) {
  println(e.Path, "has changed")
})
```

<!-- TOC -->

- [Getting Started](#getting-started)
- [API Overview](#api-overview)
- [Checks](#checks)
- [File Operations](#file-operations)
- [Directory Operations](#directory-operations)
- [Path Manipulation](#path-manipulation)
- [File System Traversal](#file-system-traversal)
- [Compression & Archiving](#compression--archiving)
- [File Hashing](#file-hashing)
- [Permissions & Ownership](#permissions--ownership)
- [Links](#links)
- [File Watching](#file-watching)
- [Path Anatomy](#path-anatomy)

<!-- /TOC -->

## Getting Started

```bash
go get github.com/renatopp/go-x
```

After installing, you can import the package and use the `fsx` name:

```go
import "github.com/renatopp/go-x/fsx"

func main() {
  fsx.Watch(context.Background(), "./assets", func (e fsx.Event) {
    checksum := fsx.ForceChecksum(e.Path)
    if fsx.IsDir(e.Path) {
      println("DIR:", checksum)
    } else {
      println("FILE:", checksum)
    }
  })
}
```

## API Overview

All functions are named to reflect how they can be used and their behavior:

| Prefix | Description | Examples |
|--------|-------------|----------|
| `*File` | Operates on files exclusively, errors on directories | `ReadFile()`, `IsFile()`, `ListFiles()` |
| `*Dir` | Operates on directories exclusively or returns directories | `EmptyDir()`, `GetHomeDir()`, `ListDirs()` |
| `*Path` | Manipulates path strings (not file system) | `JoinPath()`, `GetPathName()`, `CleanPath()` |
| `Force*` | Ignores errors and returns zero values; works with (value, error) functions | `ForceReadFile()`, `ForceSize()` |
| `*Recursive` | Operates on directories and all subdirectories | `ListFilesRecursive()`, `ChmodRecursive()` |
| `*Atomic` | Writes to temp file then renames for safety | `WriteFileAtomic()`, `WriteFileJsonAtomic()` |
| Other | Handles files and directories differently as needed | `Copy()`, `Remove()`, `Hide()` |

## Checks

Check file system state without modifying:

| Function | Description |
|----------|-------------|
| `Exists(p)` | Checks if a file or directory exists |
| `IsFile(p)` | Checks if path is a regular file |
| `IsDir(p)` | Checks if path is a directory |
| `IsEmpty(p)` | Checks if file is empty (0 bytes) or directory is empty (no entries) |
| `ForceIsEmpty(p)` | Like IsEmpty, ignores errors |
| `IsSame(p1, p2)` | Checks if two paths refer to the same file (inode comparison) |
| `IsExecutable(p)` | Checks if file has execute permission |
| `IsReadable(p)` | Checks if file is readable |
| `IsWritable(p)` | Checks if file is writable |
| `IsHidden(p)` | Checks if file/dir is hidden (starts with .) |
| `ForceIsHidden(p)` | Like IsHidden, ignores errors |
| `IsPatternValid(pattern)` | Validates a glob pattern |
| `IsAbsolutePath(p)` | Checks if path is absolute |
| `IsSlashPath(p)` | Checks if path contains forward slashes |
| `IsBackslashPath(p)` | Checks if path contains backslashes |
| `HasExtensionPath(p)` | Checks if path has a file extension |

## File Operations

Reading, writing, and modifying files:

### Reading Files

| Function | Description |
|----------|-------------|
| `OpenFile(p)` | Opens file for reading, returns *os.File |
| `ReadFile(p)` | Reads entire file as byte slice |
| `ForceReadFile(p)` | Like ReadFile, ignores errors |
| `ReadFileString(p)` | Reads entire file as string |
| `ForceReadFileString(p)` | Like ReadFileString, ignores errors |
| `ReadFileLines(p)` | Reads file and splits into lines |
| `ForceReadFileLines(p)` | Like ReadFileLines, ignores errors |
| `ReadFileJson(p, v)` | Reads JSON file and unmarshals to pointer |
| `ReadFileJsonAs[T](p)` | Type-safe JSON read with generics |
| `ForceReadFileJsonAs[T](p)` | Like ReadFileJsonAs, ignores errors |

### Writing Files

| Function | Description |
|----------|-------------|
| `CreateFile(p)` | Creates new file, returns *os.File |
| `WriteFile(p, data)` | Writes bytes to file (overwrites) |
| `WriteFileString(p, data)` | Writes string to file |
| `WriteFileLines(p, lines)` | Writes lines to file (joined by newline) |
| `WriteFileJson(p, v)` | Marshals value to JSON and writes |
| `WriteFileAtomic(p, data)` | Writes bytes atomically (temp + rename) |
| `WriteFileStringAtomic(p, data)` | Writes string atomically |
| `WriteFileLinesAtomic(p, lines)` | Writes lines atomically |
| `WriteFileJsonAtomic(p, v)` | Marshals and writes JSON atomically |
| `AppendFile(p, data)` | Appends bytes to file (creates if missing) |
| `AppendFileString(p, data)` | Appends string to file |
| `AppendFileLines(p, lines)` | Appends lines to file |
| `AppendFileJson(p, v)` | Appends JSON to file (compact, no indent) |

### File Metadata

| Function | Description |
|----------|-------------|
| `TouchFile(p)` | Creates empty file if it doesn't exist |
| `EnsureFile(p)` | Ensures file exists (creates parent dirs) |
| `TruncateFile(p, size)` | Truncates file to size bytes |
| `ReplaceInFile(p, old, new)` | Replaces all occurrences of old with new |
| `ReplaceInFileString(p, old, new)` | Like ReplaceInFile with strings |

### Temporary Files

| Function | Description |
|----------|-------------|
| `TouchTempFile(prefix)` | Creates temporary file, returns path |
| `ForceTouchTempFile(prefix)` | Like TouchTempFile, ignores errors |
| `CreateTempFile(prefix)` | Creates temporary file, returns *os.File |

## Directory Operations

Creating, managing, and inspecting directories:

### Directory Inspection

| Function | Description |
|----------|-------------|
| `ListDirs(p)` | Lists directory names in directory |
| `ForceListDirs(p)` | Like ListDirs, ignores errors |
| `ListDirsRecursive(p)` | Lists all subdirectories recursively (relative paths) |
| `ForceListDirsRecursive(p)` | Like ListDirsRecursive, ignores errors |
| `GetCurrentDir()` | Gets current working directory |
| `ForceGetCurrentDir()` | Like GetCurrentDir, ignores errors |
| `GetTempDir()` | Gets system temp directory |
| `GetCacheDir()` | Gets user cache directory |
| `ForceGetCacheDir()` | Like GetCacheDir, ignores errors |
| `GetConfigDir()` | Gets user config directory |
| `ForceGetConfigDir()` | Like GetConfigDir, ignores errors |
| `GetHomeDir()` | Gets user home directory |
| `ForceGetHomeDir()` | Like GetHomeDir, ignores errors |
| `GetParentDir(p)` | Gets parent directory of path |
| `ForceGetParentDir(p)` | Like GetParentDir, ignores errors |
| `GetParentDirName(p)` | Gets name of parent directory |
| `ForceGetParentDirName(p)` | Like GetParentDirName, ignores errors |
| `GetDirParts(p)` | Gets all path components for directory as PathParts struct |

### Directory Creation & Modification

| Function | Description |
|----------|-------------|
| `CreateDir(p)` | Creates directory with parent dirs |
| `EnsureDir(p)` | Ensures directory exists (errors on file) |
| `CreateTempDir(prefix)` | Creates temporary directory |
| `ForceCreateTempDir(prefix)` | Like CreateTempDir, ignores errors |
| `EmptyDir(p)` | Removes all contents of directory |
| `Chdir(p)` | Changes current working directory |

## Path Manipulation

String-based path operations (no file system access):

| Function | Description |
|----------|-------------|
| `JoinPath(elem...)` | Joins elements using OS separator |
| `JoinPathLinux(elem...)` | Joins using forward slashes |
| `JoinPathWindows(elem...)` | Joins using backslashes |
| `JoinPathWith(sep, elem...)` | Joins using custom separator |
| `AbsolutePath(p)` | Converts to absolute path |
| `ForceAbsolutePath(p)` | Like AbsolutePath, ignores errors |
| `RelativePath(base, target)` | Computes relative path from base to target |
| `ForceRelativePath(base, target)` | Like RelativePath, ignores errors |
| `CleanPath(p)` | Returns shortest equivalent path |
| `ToSlashPath(p)` | Converts to forward slashes |
| `FromSlashPath(p)` | Converts to OS separator |
| `ToBackslashPath(p)` | Converts to backslashes |
| `FromBackslashPath(p)` | Converts backslashes to forward slashes |
| `SplitPath(p)` | Splits path into components |
| `GetPathBase(p)` | Gets filename/dirname with extension |
| `GetPathName(p)` | Gets name without extension |
| `GetPathExtension(p)` | Gets extension with dot (.go) |
| `GetPathExtensionName(p)` | Gets extension without dot (go) |
| `GetPathParent(p)` | Gets parent directory path |
| `GetPathParentName(p)` | Gets parent directory name |
| `GetPathVolume(p)` | Gets volume name (Windows: C:, Unix: "") |
| `GetPathParts(p)` | Gets all components as PathParts struct |

## File System Traversal

Listing and walking directories:

| Function | Description |
|----------|-------------|
| `List(p)` | Lists all entries (files + dirs) in directory |
| `ForceList(p)` | Like List, ignores errors |
| `ListRecursive(p)` | Lists all entries recursively (relative paths) |
| `ForceListRecursive(p)` | Like ListRecursive, ignores errors |
| `ListFiles(p)` | Lists files only in directory |
| `ForceListFiles(p)` | Like ListFiles, ignores errors |
| `ListFilesRecursive(p)` | Lists all files recursively (relative paths) |
| `ForceListFilesRecursive(p)` | Like ListFilesRecursive, ignores errors |
| `Walk(p, fn)` | Walks directory tree calling fn for each entry |
| `Glob(dir, pattern)` | Matches files against glob pattern |
| `ForceGlob(dir, pattern)` | Like Glob, ignores errors |
| `Match(p, pattern)` | Checks if path matches glob pattern |
| `ForceMatch(p, pattern)` | Like Match, ignores errors |

## Compression & Archiving

Compressing and archiving files and directories, built entirely on the standard
library (`archive/zip`, `archive/tar`, `compress/gzip`, `compress/zlib`,
`compress/flate`, `compress/lzw`, `compress/bzip2`).

`Zip`/`Tar`/`TarGz` accept a file or directory as `src`; a directory's
contents are archived using paths relative to `src` (the directory itself is
not nested inside the archive). `Gzip`/`Zlib`/`Flate`/`Lzw` compress a single
file only and return `ErrIsDir` if `src` is a directory.

### Zip

| Function | Description |
|----------|--------------|
| `Zip(src, dest)` | Compresses a file or directory into a zip archive |
| `Unzip(src, dest)` | Extracts a zip archive into a directory |
| `OpenZip(path)` | Opens a zip archive for reading, returns `*zip.ReadCloser` |
| `CreateZip(path)` | Creates a zip archive for writing, returns `*ZipWriter` |

### Tar

| Function | Description |
|----------|--------------|
| `Tar(src, dest)` | Archives a file or directory into a tar archive |
| `Untar(src, dest)` | Extracts a tar archive into a directory |
| `OpenTar(path)` | Opens a tar archive for reading, returns `*TarReader` |
| `CreateTar(path)` | Creates a tar archive for writing, returns `*TarWriter` |

### Tar+Gzip

| Function | Description |
|----------|--------------|
| `TarGz(src, dest)` | Archives a file or directory into a gzip-compressed tar archive |
| `UntarGz(src, dest)` | Extracts a gzip-compressed tar archive into a directory |
| `OpenTarGz(path)` | Opens a gzip-compressed tar archive for reading, returns `*TarGzReader` |
| `CreateTarGz(path)` | Creates a gzip-compressed tar archive for writing, returns `*TarGzWriter` |

### Gzip

| Function | Description |
|----------|--------------|
| `Gzip(src, dest)` | Compresses a file with gzip |
| `Ungzip(src, dest)` | Decompresses a gzip-compressed file |
| `OpenGzip(path)` | Opens a gzip-compressed file for reading, returns `*GzipReader` |
| `CreateGzip(path)` | Creates a gzip-compressed file for writing, returns `*GzipWriter` |

### Zlib

| Function | Description |
|----------|--------------|
| `Zlib(src, dest)` | Compresses a file with zlib |
| `Unzlib(src, dest)` | Decompresses a zlib-compressed file |
| `OpenZlib(path)` | Opens a zlib-compressed file for reading, returns `*ZlibReader` |
| `CreateZlib(path)` | Creates a zlib-compressed file for writing, returns `*ZlibWriter` |

### Flate

Raw DEFLATE data has no header identifying it, so it cannot be auto-detected
by `Decompress`.

| Function | Description |
|----------|--------------|
| `Flate(src, dest)` | Compresses a file with raw DEFLATE |
| `Unflate(src, dest)` | Decompresses a raw DEFLATE-compressed file |
| `OpenFlate(path)` | Opens a raw DEFLATE-compressed file for reading, returns `*FlateReader` |
| `CreateFlate(path)` | Creates a raw DEFLATE-compressed file for writing, returns `*FlateWriter` |

### Lzw

LZW streams carry no header identifying them, so they cannot be auto-detected
by `Decompress`. The `*With` variants take explicit `order`/`litWidth`
parameters; the plain functions use `DefaultLzwOrder`/`DefaultLzwLitWidth`.

| Function | Description |
|----------|--------------|
| `Lzw(src, dest)` | Compresses a file with LZW using the default order/litWidth |
| `LzwWith(src, dest, order, litWidth)` | Compresses a file with LZW using explicit order/litWidth |
| `Unlzw(src, dest)` | Decompresses an LZW-compressed file using the default order/litWidth |
| `UnlzwWith(src, dest, order, litWidth)` | Decompresses an LZW-compressed file using explicit order/litWidth |
| `OpenLzw(path)` | Opens an LZW-compressed file for reading, returns `*LzwReader` |
| `OpenLzwWith(path, order, litWidth)` | Like OpenLzw with explicit order/litWidth |
| `CreateLzw(path)` | Creates an LZW-compressed file for writing, returns `*LzwWriter` |
| `CreateLzwWith(path, order, litWidth)` | Like CreateLzw with explicit order/litWidth |

### Bzip2

The standard library only implements a bzip2 reader, not a writer, so there is
no `Bzip2` compress function — only decompression is supported.

| Function | Description |
|----------|--------------|
| `Bunzip2(src, dest)` | Decompresses a bzip2-compressed file |
| `OpenBzip2(path)` | Opens a bzip2-compressed file for reading, returns `*Bzip2Reader` |

### Auto-detection

| Function | Description |
|----------|--------------|
| `Decompress(src, dest)` | Detects the format of src (zip, tar, tar.gz, gzip, zlib or bzip2) from its content and extracts/decompresses it into dest; returns `ErrUnknownFormat` otherwise |

`dest` is a directory for archive formats (zip, tar, tar.gz) and a file for
single-stream formats (gzip, zlib, bzip2), matching the corresponding `UnX`
function.

## File Hashing

Computing and verifying file/directory content hashes:

| Function | Description |
|----------|-------------|
| `MD5(p)` | Computes MD5 hash of file/directory |
| `ForceMD5(p)` | Like MD5, ignores errors |
| `SHA1(p)` | Computes SHA1 hash of file/directory |
| `ForceSHA1(p)` | Like SHA1, ignores errors |
| `SHA256(p)` | Computes SHA256 hash of file/directory |
| `ForceSHA256(p)` | Like SHA256, ignores errors |
| `Checksum(p)` | Computes checksum (MD5) of file/directory |
| `ForceChecksum(p)` | Like Checksum, ignores errors |
| `Hash(p, h)` | Computes hash using provided hash.Hash |
| `ForceHash(p, h)` | Like Hash, ignores errors |
| `Size(p)` | Gets file size (bytes) or directory size (recursive) |
| `ForceSize(p)` | Like Size, ignores errors |
| `GetModTime(p)` | Gets file modification time |
| `ForceGetModTime(p)` | Like GetModTime, ignores errors |
| `GetInfo(p)` | Gets os.FileInfo for path |
| `GetMode(p)` | Gets file permissions (os.FileMode) |

## Permissions & Ownership

Managing file permissions and ownership:

| Function | Description |
|----------|-------------|
| `SetMode(p, mode)` | Sets file permissions |
| `Chmod(p, mode)` | Alias for SetMode |
| `ChmodRecursive(p, mode)` | Sets permissions recursively |
| `ForceChmodRecursive(p, mode)` | Like ChmodRecursive, ignores errors |
| `SetModeRecursive(p, mode)` | Alias for ChmodRecursive |
| `ForceSetModeRecursive(p, mode)` | Like SetModeRecursive, ignores errors |
| `SetOwner(p, uid, gid)` | Changes file owner (uid/gid) |
| `Chown(p, uid, gid)` | Alias for SetOwner |
| `ChownRecursive(p, uid, gid)` | Changes ownership recursively |
| `ForceChownRecursive(p, uid, gid)` | Like ChownRecursive, ignores errors |
| `SetOwnerRecursive(p, uid, gid)` | Alias for ChownRecursive |
| `ForceSetOwnerRecursive(p, uid, gid)` | Like SetOwnerRecursive, ignores errors |
| `SetHidden(p, hidden)` | Hides/unhides file by renaming (Unix) |
| `Hide(p)` | Hides file/directory |
| `Unhide(p)` | Unhides file/directory |

## Links

Creating and managing symbolic/hard links:

| Function | Description |
|----------|-------------|
| `Link(src, dst)` | Creates hard link from src to dst |
| `Symlink(oldname, newname)` | Creates symbolic link |
| `Readlink(p)` | Reads destination of symbolic link |
| `ForceReadlink(p)` | Like Readlink, ignores errors |

## File System Operations

General file system operations:

| Function | Description |
|----------|-------------|
| `Copy(src, dst)` | Recursively copies file or directory |
| `Move(src, dst)` | Moves/renames file or directory |
| `Rename(old, new)` | Renames file or directory |
| `Remove(p)` | Removes file or directory (recursive) |
| `Empty(p)` | Empties file (truncates) or directory (removes contents) |

## File Watching

Monitoring file system changes:

| Function | Description |
|----------|-------------|
| `NewWatcher()` | Creates new file system watcher |
| `Watch(ctx, p, callback)` | Watches single path for changes |
| `WatchRecursive(ctx, p, callback)` | Watches directory and subdirectories |
| `WatchGlob(ctx, dir, pattern, callback)` | Watches with glob pattern filtering |

### Watcher Methods

```go
watcher, err := fsx.NewWatcher()
if err != nil {
  panic(err)
}
defer watcher.Close()

watcher.Add(path)          // Add path to watch list
watcher.Remove(path)       // Remove path from watch list
watcher.Has(path)          // Check if path is being watched
watcher.WatchList()        // Get list of watched paths
watcher.Watch(ctx, callback) // Start watching
```

### Watch Events

Events contain:
- `Op` - Operation type (bitmasked), check with `event.Has(fsx.EvtCreate)`:
  - `EvtCreate` - File/directory created
  - `EvtWrite` - File written
  - `EvtRemove` - File/directory removed
  - `EvtRename` - File/directory renamed
  - `EvtChmod` - File permissions changed
  - `EvtError` - Error occurred
- `Path` - Full path of file/directory that changed
- `Err` - Error if Op contains EvtError

## Path Anatomy

You can use `GetPathParts` or `GetDirParts` to extract all path information at once.

For the path `/c/users/dev/fs/path.go`:

| Part | Example | Description |
|------|---------|-------------|
| **Absolute** | `/c/users/dev/fs/path.go` | Complete absolute path |
| **Base** | `path.go` | File/directory name with extension |
| **Name** | `path` | File/directory name without extension |
| **Ext** | `.go` | Extension including dot |
| **ExtName** | `go` | Extension without dot |
| **Parent** | `/c/users/dev/fs` | Parent directory path |
| **ParentName** | `fs` | Parent directory name |
| **Volume** | `c:` (Windows) / `` (Unix) | Drive letter (Windows only) |

### PathParts Struct

```go
type PathParts struct {
  Absolute   string // /c/users/dev/fs/path.go
  Base       string // path.go
  Name       string // path
  Ext        string // .go
  ExtName    string // go
  Parent     string // /c/users/dev/fs
  ParentName string // fs
  Volume     string // c: (Windows), "" (Unix)
}

// Get all parts at once
parts := fsx.GetPathParts("./path/to/file.go")
// or for directories
parts := fsx.GetDirParts("./path/to/dir")
```
