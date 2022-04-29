package safe

import (
	"bytes"
	"strings"

	"golang.org/x/exp/slices"
)

const nullbyte = string(rune(0)) // or "\x00"

// Contains reports whether nullbyte is within s.
func ContainsNullByte(a any) (result bool, ok bool) {
	switch v := a.(type) {
	case string:
		return strings.Contains(v, nullbyte), true
	case []byte:
		return slices.Contains(v, 0), true
	}
	return false, false
}

// TrimNullByte returns a string or bytes that is trimmed null bytes.
func TrimNullByte(a any) (result any, ok bool) {
	switch v := a.(type) {
	case string:
		return strings.Trim(v, nullbyte), true
	case []byte:
		return bytes.Trim(v, nullbyte), true
	}
	return a, false
}

// RemoveNullByte returns a string or bytes that is removed null bytes.
func RemoveNullByte(a any) (result any, ok bool) {
	switch v := a.(type) {
	case string:
		return strings.ReplaceAll(v, nullbyte, ""), true
	case []byte:
		return bytes.ReplaceAll(v, []byte(nullbyte), []byte("")), true
	}
	return a, false
}
