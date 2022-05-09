package nullbyte_test

import (
	"testing"

	"github.com/exaream/go-lib/security/nullbyte"
	"github.com/google/go-cmp/cmp"
)

const (
	nullByteStr   = "\x00"
	noNullbyteStr = "foobar"
	before        = nullByteStr + "foo" + nullByteStr + "bar" + nullByteStr
	afterTrim     = "foo" + nullByteStr + "bar"
)

func TestContainsInStr(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		in string

		want bool
	}{
		"null byte":    {before, true},
		"no null byte": {noNullbyteStr, false},
		"empty":        {"", false},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := nullbyte.Contains(tt.in)
			if tt.want != got {
				t.Errorf("in=%s, want=%t, got=%t", tt.in, tt.want, got)
			}
		})
	}
}

func TestContainsInByte(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		in []byte

		want bool
	}{
		"null byte":    {[]byte(before), true},
		"no null byte": {[]byte(noNullbyteStr), false},
		"empty":        {[]byte(""), false},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := nullbyte.Contains(tt.in)
			if tt.want != got {
				t.Errorf("in=%s, want=%t, got=%t", tt.in, tt.want, got)
			}
		})
	}
}

func TestTrimFromStr(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		in string

		want string
	}{
		"null byte":    {before, afterTrim},
		"no null byte": {noNullbyteStr, noNullbyteStr},
		"empty":        {"", ""},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := nullbyte.Trim(tt.in)
			if tt.want != got {
				t.Errorf("in=%s, want=%s, got=%s", tt.in, tt.want, got)
			}
		})
	}
}

func TestTrimFromByte(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		in []byte

		want []byte
	}{
		"null byte":    {[]byte(before), []byte(afterTrim)},
		"no null byte": {[]byte(noNullbyteStr), []byte(noNullbyteStr)},
		"empty":        {[]byte(""), []byte("")},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := nullbyte.Trim(tt.in)
			if !cmp.Equal(tt.want, got) {
				t.Errorf("in=%v, want=%v, got=%v", tt.in, tt.want, got)
			}
		})
	}
}

func TestRemoveAllFromStr(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		in string

		want string
	}{
		"null byte":    {before, noNullbyteStr},
		"no null byte": {noNullbyteStr, noNullbyteStr},
		"empty":        {"", ""},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := nullbyte.RemoveAll(tt.in)
			if tt.want != got {
				t.Errorf("in=%s, want=%s, got=%s", tt.in, tt.want, got)
			}
		})
	}
}

func TestRemoveAllFromByte(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		in []byte

		want []byte
	}{
		"null byte":    {[]byte(before), []byte(noNullbyteStr)},
		"no null byte": {[]byte(noNullbyteStr), []byte(noNullbyteStr)},
		"empty":        {[]byte(""), nil},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := nullbyte.RemoveAll(tt.in)
			if !cmp.Equal(tt.want, got) {
				t.Errorf("in=%v, want=%v, got=%v", tt.in, tt.want, got)
			}
		})
	}
}
