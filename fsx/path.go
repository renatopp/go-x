package fsx

import (
	"path/filepath"
	"strings"
)

type PathParts struct {
	Absolute   string // Absolute path (eg: /home/users/dev/fs/path.go)
	Base       string // Base name (eg: path.go)
	Name       string // Name without extension (eg: path)
	Ext        string // Extension with dot (eg: .go)
	ExtName    string // Extension without dot (eg: go)
	Parent     string // Parent directory (eg: /home/users/dev/fs)
	ParentName string // Parent directory name (eg: fs)
	Volume     string // Volume name (eg: C: on Windows, empty on Unix)
}

// JoinPath joins path elements using the OS-specific path separator.
func JoinPath(elem ...string) string {
	return filepath.Join(elem...)
}

// JoinPathLinux joins path elements using forward slashes (Unix-style).
func JoinPathLinux(elem ...string) string {
	return strings.Join(elem, "/")
}

// JoinPathWindows joins path elements using backslashes (Windows-style).
func JoinPathWindows(elem ...string) string {
	return strings.Join(elem, "\\")
}

// JoinPathWith joins path elements using a specific separator.
func JoinPathWith(sep string, elem ...string) string {
	return strings.Join(elem, sep)
}

// AbsolutePath converts a path to an absolute path.
func AbsolutePath(p string) (string, error) {
	absPath, err := filepath.Abs(p)
	if err != nil {
		return "", err
	}
	return absPath, nil
}

// ForceAbsolutePath is like AbsolutePath but ignores errors and returns an empty string on error.
func ForceAbsolutePath(p string) string {
	abs, _ := AbsolutePath(p)
	return abs
}

// RelativePath returns the relative path from base to target.
func RelativePath(base, target string) (string, error) {
	absBase, err := AbsolutePath(base)
	if err != nil {
		return target, err
	}
	absTarget, err := AbsolutePath(target)
	if err != nil {
		return target, err
	}
	relPath, err := filepath.Rel(absBase, absTarget)
	if err != nil {
		return target, err
	}
	return relPath, nil
}

// ForceRelativePath is like RelativePath but ignores errors and returns the target path on error.
func ForceRelativePath(base, target string) string {
	rel, _ := RelativePath(base, target)
	return rel
}

// IsAbsolutePath checks if a path is absolute.
func IsAbsolutePath(p string) bool {
	return filepath.IsAbs(p)
}

// CleanPath returns the shortest path equivalent to p by eliminating . and .. elements.
func CleanPath(p string) string {
	return filepath.Clean(p)
}

// ToBackslashPath converts forward slashes to backslashes.
func ToBackslashPath(p string) string {
	return strings.ReplaceAll(p, "/", "\\")
}

// FromBackslashPath converts backslashes to forward slashes.
func FromBackslashPath(p string) string {
	return strings.ReplaceAll(p, "\\", "/")
}

// ToSlashPath converts the path to use forward slashes.
func ToSlashPath(p string) string {
	return filepath.ToSlash(p)
}

// FromSlashPath converts forward slashes to the OS-specific separator.
func FromSlashPath(p string) string {
	return filepath.FromSlash(p)
}

// IsSlashPath checks if a path contains forward slashes.
func IsSlashPath(p string) bool {
	return strings.Contains(p, "/")
}

// IsBackslashPath checks if a path contains backslashes.
func IsBackslashPath(p string) bool {
	return strings.Contains(p, "\\")
}

// HasExtensionPath checks if a path has a file extension.
func HasExtensionPath(p string) bool {
	return filepath.Ext(p) != ""
}

// SplitPath splits a path into its components using forward slashes.
func SplitPath(p string) []string {
	return strings.Split(ToSlashPath(p), "/")
}

// GetPathBase returns the last element of the path, which is typically the
// file name or the last directory in the path.
//
//	/home/users/dev/fs/path.go -> path.go
//	/home/users/dev/fs/ -> fs
//	/home/users/dev/fs -> fs
func GetPathBase(p string) string {
	return filepath.Base(p)
}

// GetPathName returns the file name without the extension.
//
//	/home/users/dev/fs/path.go -> path
//	/home/users/dev/fs/ -> fs
//	/home/users/dev/fs -> fs
func GetPathName(p string) string {
	ext := filepath.Ext(p)
	return strings.TrimSuffix(filepath.Base(p), ext)
}

// GetPathExtension returns the file extension, including the dot.
//
//	/home/users/dev/fs/path.go -> .go
//	/home/users/dev/fs/ -> ""
//	/home/users/dev/fs -> ""
func GetPathExtension(p string) string {
	return filepath.Ext(p)
}

// GetPathExtensionName returns the file extension without the dot.
//
//	/home/users/dev/fs/path.go -> go
//	/home/users/dev/fs/ -> ""
//	/home/users/dev/fs -> ""
func GetPathExtensionName(p string) string {
	ext := filepath.Ext(p)
	return strings.TrimPrefix(ext, ".")
}

// GetPathParent returns the parent directory of the given path.
//
//	/home/users/dev/fs/path.go -> /home/users/dev/fs
//	/home/users/dev/fs/ -> /home/users/dev
//	/home/users/dev/fs -> /home/users/dev
func GetPathParent(p string) string {
	return filepath.Dir(p)
}

// GetPathParentName returns the name of the parent directory of the given path.
//
//	/home/users/dev/fs/path.go -> fs
//	/home/users/dev/fs/ -> dev
//	/home/users/dev/fs -> dev
func GetPathParentName(p string) string {
	dir := filepath.Dir(p)
	return filepath.Base(dir)
}

// GetPathVolume returns the volume name of the given path. On Windows, this is
// the drive letter (e.g., "C:"). On Unix-like systems, this will be an empty string.
//
//	C:\Users\dev\fs\path.go -> C:
func GetPathVolume(p string) string {
	return filepath.VolumeName(p)
}

// GetPathParts returns a PathParts struct containing various components of
// the given path.
func GetPathParts(p string) PathParts {
	abs := ForceAbsolutePath(p)
	return PathParts{
		Absolute:   abs,
		Base:       GetPathBase(abs),
		Name:       GetPathName(abs),
		Ext:        GetPathExtension(abs),
		ExtName:    GetPathExtensionName(abs),
		Parent:     GetPathParent(abs),
		ParentName: GetPathParentName(abs),
		Volume:     GetPathVolume(abs),
	}
}
