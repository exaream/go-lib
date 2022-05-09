package nullbyte

import (
	"bytes"
	"strings"

	"golang.org/x/exp/slices"
)

const nullByteStr = string(rune(0)) // "\x00"

type Sequence interface {
	string | []byte
}

// Contains reports whether null byte is within a.
func Contains[T Sequence](a T) bool {
	switch v := any(a).(type) {
	case string:
		return strings.Contains(v, nullByteStr)
	case []byte:
		return slices.Contains(v, 0)
	}
	// We do not pass this.
	return false
}

// Trim returns a string or bytes that is trimmed null bytes.
func Trim[T Sequence](a T) any {
	switch v := any(a).(type) {
	case string:
		return strings.Trim(v, nullByteStr)
	case []byte:
		return bytes.Trim(v, nullByteStr)
	}
	return a
}

// RemoveAll returns a string or bytes that is removed null bytes.
func RemoveAll[T Sequence](a T) any {
	switch v := any(a).(type) {
	case string:
		return strings.ReplaceAll(v, nullByteStr, "")
	case []byte:
		return bytes.ReplaceAll(v, []byte(nullByteStr), []byte(""))
	}
	return a
}
