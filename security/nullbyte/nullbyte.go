package nullbyte

import (
	"bytes"
	"strings"

	"golang.org/x/exp/slices"
)

const nullByteStr = string(rune(0)) // "\x00"

// Interface of arguments and return values of functions.
type Sequence interface {
	string | []byte
}

// Contains reports whether null byte is within a.
func Contains[T Sequence](a T) bool {
	// Use any() to do type assertion.
	switch v := any(a).(type) {
	case string:
		return strings.Contains(v, nullByteStr)
	case []byte:
		return slices.Contains(v, 0)
	}
	// We never pass through here..
	return false
}

// Trim returns a string or bytes that is trimmed null bytes.
func Trim[T Sequence](a T) T {
	// Use any() to do type assertion.
	switch v := any(a).(type) {
	// Do type assertion again to return as T.
	case string:
		return any(strings.Trim(v, nullByteStr)).(T)
	case []byte:
		return any(bytes.Trim(v, nullByteStr)).(T)
	}
	return a
}

// RemoveAll returns a string or bytes that is removed null bytes.
func RemoveAll[T Sequence](a T) T {
	// Use any() to do type assertion.
	switch v := any(a).(type) {
	// Do type assertion again to return as T.
	case string:
		return any(strings.ReplaceAll(v, nullByteStr, "")).(T)
	case []byte:
		return any(bytes.ReplaceAll(v, []byte(nullByteStr), []byte(""))).(T)
	}
	return a
}
