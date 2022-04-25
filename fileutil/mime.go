// https://docs.w3cub.com/http/basics_of_http/mime_types/complete_list_of_mime_types.html
package fileutil

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// MimeType returns MimeType by a file name.
func MimeType(file string) string {
	ext := filepath.Ext(file)
	return MimeTypeByExt(ext)
}

// MimeType returns MimeType by an extension.
func MimeTypeByExt(ext string) string {
	if strings.HasPrefix(ext, ".") {
		ext = strings.TrimLeft(ext, ".")
	}

	ext = strings.ToLower(ext)

	if typ, ok := mimeTypes[ext]; ok {
		return typ
	}

	return ""
}

// ContentType returns Content-Type by a file name.
func ContentType(file string) (string, error) {
	f, err := os.Open(filepath.Clean(file))

	if err != nil {
		return "", err
	}

	buf := make([]byte, 512)

	if _, err := f.Read(buf); err != nil {
		return "", err
	}

	typ := http.DetectContentType(buf)

	if _, err := f.Seek(0, 0); err != nil {
		return "", err
	}

	return typ, nil
}
