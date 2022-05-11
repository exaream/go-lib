// html package's extension
package htmlex

import (
	"html"
	"strings"
)

func EscapeString(s string) string {
	s = html.EscapeString(s)
	for char, ref := range specialChars {
		s = strings.ReplaceAll(s, char, ref)
	}
	return s
}

func UnescapeString(s string) string {
	s = html.UnescapeString(s)
	for char, ref := range specialChars {
		s = strings.ReplaceAll(s, ref, char)
	}
	return s
}
