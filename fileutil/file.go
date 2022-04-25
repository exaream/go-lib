package fileutil

import (
	"os"
	"path/filepath"
	"strings"
)

// Exist returns true if a directory of a file exists.
func Exist(target string) bool {
	if _, err := os.Stat(target); err != nil {
		return false
	}
	return true
}

// Ext returns an extension of a file.
func Ext(filePath string) string {
	return strings.ToLower(filepath.Ext(filePath))
}

// Stem returns a file's stem (a file name without the extension.) by a path
func Stem(file string) string {
	base := filepath.Base(file)
	ext := filepath.Ext(base)
	return strings.TrimRight(base, ext) // or base[:len(base)-len(ext)]
}
