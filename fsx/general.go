package fsx

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"hash"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/bmatcuk/doublestar/v4"
)

// ----------------------------------------------------------------------------
// CHECKS
// ----------------------------------------------------------------------------

// Exists checks if a file or directory exists at the given p.
func Exists(p string) bool {
	_, err := os.Stat(p)
	return !os.IsNotExist(err)
}

// IsEmpty checks if a file or directory at the specified path is empty.
// For files, it checks if the file size is zero bytes.
// For directories, it checks if the directory contains no files or subdirectories.
// If the path does not exist, it returns an error but also true.
func IsEmpty(p string) (bool, error) {
	if IsDir(p) {
		return isEmptyDir(p)
	} else if IsFile(p) {
		return isEmptyFile(p)
	} else {
		return true, os.ErrNotExist
	}
}

// ForceIsEmpty is like IsEmpty but ignores errors and returns false on error.
func ForceIsEmpty(p string) bool {
	empty, _ := IsEmpty(p)
	return empty
}

// IsSame checks if two paths refer to the same file by comparing their inode information.
func IsSame(p1, p2 string) bool {
	s1, err := os.Stat(p1)
	if err != nil {
		return false
	}
	s2, err := os.Stat(p2)
	if err != nil {
		return false
	}
	return os.SameFile(s1, s2)
}

// IsExecutable checks if a file at the specified path is executable.
func IsExecutable(p string) bool {
	if !IsFile(p) {
		return false
	}
	info, err := os.Stat(p)
	if err != nil {
		return false
	}
	mode := info.Mode()
	if mode&0111 != 0 {
		return true
	}
	if runtime.GOOS == "windows" {
		ext := strings.ToLower(filepath.Ext(p))
		pathext := os.Getenv("PATHEXT")
		for _, e := range strings.Split(pathext, ";") {
			if strings.ToLower(e) == ext {
				return true
			}
		}
	}
	return false
}

// IsReadable checks if a file at the specified path is readable.
func IsReadable(p string) bool {
	if !IsFile(p) {
		return false
	}
	file, err := os.OpenFile(p, os.O_RDONLY, 0)
	if err != nil {
		return false
	}
	file.Close()
	return true
}

// IsWritable checks if a file at the specified path is writable.
func IsWritable(p string) bool {
	if !IsFile(p) {
		return false
	}
	file, err := os.OpenFile(p, os.O_WRONLY, 0)
	if err != nil {
		return false
	}
	file.Close()
	return true
}

// IsHidden checks if a file or directory is hidden (starts with a dot on Unix-like systems).
func IsHidden(p string) (bool, error) {
	// abs := Force(AbsolutePath(p))
	base := filepath.Base(p)
	// if runtime.GOOS == "windows" {
	// 	pointer, err := syscall.UTF16PtrFromString(abs)
	// 	if err != nil {
	// 		return false, err
	// 	}
	// 	attributes, err := syscall.GetFileAttributes(pointer)
	// 	if err != nil {
	// 		return false, err
	// 	}
	// 	return attributes&syscall.FILE_ATTRIBUTE_HIDDEN != 0, nil
	// }
	return strings.HasPrefix(base, "."), nil
}

// IsPatternValid checks if the given glob pattern is valid.
func IsPatternValid(pattern string) bool {
	return doublestar.ValidatePattern(pattern)
}

// ForceIsHidden is like IsHidden but ignores errors and returns false on error.
func ForceIsHidden(p string) bool {
	hidden, _ := IsHidden(p)
	return hidden
}

// ----------------------------------------------------------------------------
// TRAVERSAL
// ----------------------------------------------------------------------------

// Walk traverses the directory tree rooted at the specified path, calling the
// provided function for each file or directory encountered. The function receives
// the full path of the file or directory as its argument. If the function
// returns an error, the walk is aborted and the error is returned.
func Walk(p string, fn func(string) error) error {
	return filepath.WalkDir(p, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		return fn(path)
	})
}

// List returns a slice of names of all entries (files and directories) within
// the specified directory path. If the directory does not exist or is not
// accessible, it returns an error. This function does not include the full paths,
// only the names of the entries.
//
// This function is not recursive; it only lists entries in the specified
// directory, not in its subdirectories.
func List(p string) ([]string, error) {
	entries, err := os.ReadDir(p)
	files := []string{}
	if err != nil {
		return files, err
	}
	names := make([]string, len(entries))
	for i, entry := range entries {
		names[i] = entry.Name()
	}
	return names, nil
}

// ForceList is like List but ignores errors and returns an empty slice on error.
func ForceList(p string) []string {
	list, _ := List(p)
	return list
}

// ListRecursive returns a slice of relative paths of all entries (files and
// directories) within the specified directory path and its subdirectories.
// If the directory does not exist or is not accessible, it returns an error.
// The returned paths are relative to the specified directory.
//
// This function is recursive; it lists entries in the specified directory
// and all its subdirectories.
func ListRecursive(p string) ([]string, error) {
	if !IsDir(p) {
		return nil, ErrNotDir
	}

	results := []string{}
	err := filepath.WalkDir(p, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		relPath, err := filepath.Rel(p, path)
		if err != nil {
			return err
		}
		if relPath != "." {
			results = append(results, relPath)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return results, nil
}

// ForceListRecursive is like ListRecursive but ignores errors and returns an empty slice on error.
func ForceListRecursive(p string) []string {
	list, _ := ListRecursive(p)
	return list
}

// Glob returns the names of all files matching pattern or nil if there is no
// matching file. The syntax of patterns is the same as in filepath.Match.
// The pattern may describe hierarchical names such as /usr/*/bin/ed (assuming
// the Separator is '/').
func Glob(dir, pattern string) ([]string, error) {
	files, err := doublestar.Glob(os.DirFS("."), filepath.Join(dir, pattern))
	if files == nil {
		files = []string{}
	}
	r := len(dir) + len(PathSeparator)
	for i, f := range files {
		files[i] = f[r:]
	}
	return files, err
}

// ForceGlob is like Glob but ignores errors and returns an empty slice on error.
func ForceGlob(dir string, pattern string) []string {
	files, _ := Glob(dir, pattern)
	return files
}

// Match checks if a path matches the given glob pattern.
func Match(p, pattern string) (bool, error) {
	return doublestar.Match(pattern, p)
}

// ForceMatch is like Match but ignores errors and returns false on error.
func ForceMatch(p, pattern string) bool {
	matched, _ := Match(p, pattern)
	return matched
}

// ----------------------------------------------------------------------------
// OPERATIONS
// ----------------------------------------------------------------------------

// Copy copies a file or directory from src to dst. If src is a directory, it
// copies the entire directory recursively. If src is a file, it copies the file.
// If dst does not exist, it will be created. If it exists, it will be merged
// (for directories) or overwritten (for files).
func Copy(src, dst string) error {
	if IsDir(src) {
		return copyDir(src, dst)
	}
	return copyFile(src, dst)
}

// Move moves a file or directory from src to dst. It is equivalent to renaming
// the file or directory. If src and dst are on different filesystems, it
// performs a copy followed by a delete of the original.
func Move(src, dst string) error {
	return os.Rename(src, dst)
}

// Rename renames (moves) a file or directory from oldPath to newPath. If oldPath
// and newPath are on different filesystems, it performs a copy followed by a
// delete of the original.
func Rename(oldPath, newPath string) error {
	return os.Rename(oldPath, newPath)
}

// Remove removes a file or directory at the specified path. If the path is a
// directory, it removes the directory and all its contents recursively.
// If the path does not exist, it returns nil (no error).
// If there is an error, it will be of type [*PathError].
func Remove(p string) error {
	return os.RemoveAll(p)
}

// SetMode sets the file mode (permissions) of a file at the specified path. If
// the path does not exist, it returns an error.
func SetMode(p string, mode os.FileMode) error {
	return os.Chmod(p, mode)
}

// SetModeRecursive recursively sets the file mode (permissions) for a directory and all its contents.
func SetModeRecursive(p string, mode os.FileMode) error {
	return ChmodRecursive(p, mode)
}

// ForceSetModeRecursive is like SetModeRecursive but ignores errors.
func ForceSetModeRecursive(p string, mode os.FileMode) {
	ForceChmodRecursive(p, mode)
}

// SetHidden sets the hidden status of a file or directory by renaming it to/from a dot-prefixed name.
func SetHidden(p string, hidden bool) error {
	// abs := Force(AbsolutePath(p))
	if runtime.GOOS == "windows" {
		// pointer, err := syscall.UTF16PtrFromString(abs)
		// if err != nil {
		// 	return err
		// }
		// attributes, err := syscall.GetFileAttributes(pointer)
		// if err != nil {
		// 	return err
		// }
		// if hidden {
		// 	attributes |= syscall.FILE_ATTRIBUTE_HIDDEN
		// } else {
		// 	attributes &^= syscall.FILE_ATTRIBUTE_HIDDEN
		// }
		// return syscall.SetFileAttributes(pointer, attributes)
	}
	base := filepath.Base(p)
	dir := filepath.Dir(p)
	if hidden {
		if strings.HasPrefix(base, ".") {
			return nil
		}
		newPath := filepath.Join(dir, "."+base)
		return os.Rename(p, newPath)
	} else {
		if !strings.HasPrefix(base, ".") {
			return nil
		}
		newBase := strings.TrimPrefix(base, ".")
		newPath := filepath.Join(dir, newBase)
		return os.Rename(p, newPath)
	}
}

// Hide hides a file or directory by renaming it to a dot-prefixed name.
func Hide(p string) error {
	return SetHidden(p, true)
}

// Unhide unhides a file or directory by removing its dot prefix.
func Unhide(p string) error {
	return SetHidden(p, false)
}

// Chmod is an alias for SetMode.
func Chmod(p string, mode os.FileMode) error {
	return os.Chmod(p, mode)
}

// ChmodRecursive recursively changes permissions for a directory and all its contents.
func ChmodRecursive(p string, mode os.FileMode) error {
	if !IsDir(p) {
		return os.Chmod(p, mode)
	}
	err := os.Chmod(p, mode)
	if err != nil {
		return err
	}
	Walk(p, func(path string) error {
		return os.Chmod(path, mode)
	})
	return nil
}

// ForceChmodRecursive is like ChmodRecursive but ignores errors.
func ForceChmodRecursive(p string, mode os.FileMode) {
	if !IsDir(p) {
		os.Chmod(p, mode)
		return
	}
	os.Chmod(p, mode)
	Walk(p, func(path string) error {
		os.Chmod(path, mode)
		return nil
	})
}

// SetOwner changes the ownership of a file (alias for Chown).
func SetOwner(p string, uid, gid int) error {
	return os.Chown(p, uid, gid)
}

// SetOwnerRecursive recursively changes ownership for a directory and all its contents.
func SetOwnerRecursive(p string, uid, gid int) error {
	return ChownRecursive(p, uid, gid)
}

// ForceSetOwnerRecursive is like SetOwnerRecursive but ignores errors.
func ForceSetOwnerRecursive(p string, uid, gid int) {
	ForceChownRecursive(p, uid, gid)
}

// Chown changes the ownership of a file at the specified path to the given
// user ID (uid) and group ID (gid). If the path does not exist, it returns
// an error.
func Chown(p string, uid, gid int) error {
	return os.Chown(p, uid, gid)
}

// ChownRecursive recursively changes ownership for a directory and all its contents.
func ChownRecursive(p string, uid, gid int) error {
	if !IsDir(p) {
		return os.Chown(p, uid, gid)
	}
	err := os.Chown(p, uid, gid)
	if err != nil {
		return err
	}
	Walk(p, func(path string) error {
		return os.Chown(path, uid, gid)
	})
	return nil
}

// ForceChownRecursive is like ChownRecursive but ignores errors.
func ForceChownRecursive(p string, uid, gid int) {
	if !IsDir(p) {
		os.Chown(p, uid, gid)
		return
	}
	os.Chown(p, uid, gid)
	Walk(p, func(path string) error {
		os.Chown(path, uid, gid)
		return nil
	})
}

// Chdir changes the current working directory to the specified path. If the
// path does not exist or is not a directory, it returns an error.
func Chdir(p string) error {
	return os.Chdir(p)
}

// Empty empties a file (truncates to zero) or directory (removes all contents).
func Empty(p string) error {
	if IsDir(p) {
		return EmptyDir(p)
	} else if IsFile(p) {
		return TruncateFile(p, 0)
	} else {
		return os.ErrNotExist
	}
}

// ----------------------------------------------------------------------------
// LINKS
// ----------------------------------------------------------------------------

// Link creates a hard link from src to dst. If src does not exist or dst
// already exists, it returns an error.
func Link(src, dst string) error {
	return os.Link(src, dst)
}

// Symlink creates a symbolic link from oldname to newname. If oldname does not
// exist or newname already exists, it returns an error.
func Symlink(oldname, newname string) error {
	return os.Symlink(oldname, newname)
}

// Readlink returns the destination of the named symbolic link.
func Readlink(p string) (string, error) {
	return os.Readlink(p)
}

// ForceReadlink is like Readlink but ignores errors and returns an empty string on error.
func ForceReadlink(p string) string {
	link, _ := Readlink(p)
	return link
}

// ----------------------------------------------------------------------------
// HASHING
// ----------------------------------------------------------------------------

// MD5 computes the MD5 hash of a file or directory.
func MD5(p string) (string, error) {
	return Hash(p, md5.New())
}

// ForceMD5 is like MD5 but ignores errors and returns an empty string on error.
func ForceMD5(p string) string {
	sum, _ := MD5(p)
	return sum
}

// SHA1 computes the SHA1 hash of a file or directory.
func SHA1(p string) (string, error) {
	return Hash(p, sha1.New())
}

// ForceSHA1 is like SHA1 but ignores errors and returns an empty string on error.
func ForceSHA1(p string) string {
	sum, _ := SHA1(p)
	return sum
}

// SHA256 computes the SHA256 hash of a file or directory.
func SHA256(path string) (string, error) {
	return Hash(path, sha256.New())
}

// ForceSHA256 is like SHA256 but ignores errors and returns an empty string on error.
func ForceSHA256(path string) string {
	sum, _ := SHA256(path)
	return sum
}

// Checksum computes the checksum (MD5) of a file or directory.
func Checksum(p string) (string, error) {
	return Hash(p, md5.New())
}

// ForceChecksum is like Checksum but ignores errors and returns an empty string on error.
func ForceChecksum(p string) string {
	sum, _ := Checksum(p)
	return sum
}

// Hash computes the hash of a file or directory using the provided hash function.
func Hash(p string, h hash.Hash) (string, error) {
	if IsDir(p) {
		return hashDir(p, h)
	}
	return hashFile(p, h)
}

// ForceHash is like Hash but ignores errors and returns an empty string on error.
func ForceHash(p string, h hash.Hash) string {
	sum, _ := Hash(p, h)
	return sum
}

// ----------------------------------------------------------------------------
// INFO
// ----------------------------------------------------------------------------

// Size returns the size of a file or directory at the specified path in bytes.
// If the path is a directory, it computes the total size of all files within
// the directory recursively. It returns the size in bytes.
func Size(p string) (int64, error) {
	if IsDir(p) {
		return sizeDir(p)
	}
	return sizeFile(p)
}

// ForceSize is like Size but ignores errors and returns 0 on error.
func ForceSize(p string) int64 {
	size, _ := Size(p)
	return size
}

// GetModTime returns the modification time of a file at the specified path as a
// Unix timestamp (seconds since January 1, 1970). If the path does not exist
// or is a directory, it returns an error.
func GetModTime(p string) (time.Time, error) {
	info, err := os.Stat(p)
	if err != nil {
		return time.Time{}, err
	}
	return info.ModTime(), nil
}

// ForceGetModTime is like GetModTime but ignores errors and returns zero time on error.
func ForceGetModTime(p string) time.Time {
	t, _ := GetModTime(p)
	return t
}

// GetInfo returns a FileInfo describing the file at the specified path. If the
// path does not exist, it returns an error.
func GetInfo(p string) (os.FileInfo, error) {
	return os.Stat(p)
}

// GetMode returns the file mode (permissions) of a file at the specified path. If
// the path does not exist, it returns an error.
func GetMode(p string) (os.FileMode, error) {
	info, err := os.Stat(p)
	if err != nil {
		return 0, err
	}
	return info.Mode(), nil
}

// Getwd returns the current working directory.
func GetPwd() (string, error) {
	return GetCurrentDir()
}

// Pwd is an alias for Getwd.
func Pwd() (string, error) {
	return GetCurrentDir()
}
