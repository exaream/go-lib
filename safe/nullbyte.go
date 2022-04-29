package safe

import (
	"strings"

	"golang.org/x/exp/slices"
)

const nullbyte = string(rune(0)) // or "\x00"

// Contains reports whether nullbyte is within s
func Contains(i any) (result, ok bool) {
	switch v := i.(type) {
	case string:
		return strings.Contains(string(v), nullbyte), true
	case []byte:
		return slices.Contains(v, 0), true
	}
	return false, false
}
