package fileutil

import (
	"os"
	"path/filepath"
	"strings"
)

func Exist(target string) bool {
	if _, err := os.Stat(target); err != nil {
		return false
	}
	return true
}

func Ext(filePath string) string {
	return strings.ToLower(filepath.Ext(filePath))
}
