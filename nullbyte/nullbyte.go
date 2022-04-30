package nullbyte

import (
	"bytes"
	"strings"

	"golang.org/x/exp/slices"
)

const nullByteStr = string(rune(0)) // "\x00"

// Contains reports whether null byte is within a.
func Contains(a any) (result bool, ok bool) {
	switch v := a.(type) {
	case string:
		return strings.Contains(v, nullByteStr), true
	case []byte:
		return slices.Contains(v, 0), true
	}
	return false, false
}

// Trim returns a string or bytes that is trimmed null bytes.
func Trim(a any) (result any, ok bool) {
	switch v := a.(type) {
	case string:
		return strings.Trim(v, nullByteStr), true
	case []byte:
		return bytes.Trim(v, nullByteStr), true
	}
	return a, false
}

// RemoveAll returns a string or bytes that is removed null bytes.
func RemoveAll(a any) (result any, ok bool) {
	switch v := a.(type) {
	case string:
		return strings.ReplaceAll(v, nullByteStr, ""), true
	case []byte:
		return bytes.ReplaceAll(v, []byte(nullByteStr), []byte("")), true
	}
	return a, false
}
