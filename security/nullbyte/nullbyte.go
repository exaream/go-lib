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
	// We do not pass here.
	return false
}

// Trim returns a string or bytes that is trimmed null bytes.
func Trim[T Sequence](a T) T {
	switch v := any(a).(type) {
	case string:
		return any(strings.Trim(v, nullByteStr)).(T)
	case []byte:
		return any(bytes.Trim(v, nullByteStr)).(T)
	}
	return a
}

// RemoveAll returns a string or bytes that is removed null bytes.
func RemoveAll[T Sequence](a T) T {
	switch v := any(a).(type) {
	case string:
		return any(strings.ReplaceAll(v, nullByteStr, "")).(T)
	case []byte:
		return any(bytes.ReplaceAll(v, []byte(nullByteStr), []byte(""))).(T)
	}
	return a
}
