package nullbyte_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/exaream/go-lib/safe/nullbyte"
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

		want   bool
		wantOK bool
	}{
		"null byte":    {before, true, true},
		"no null byte": {noNullbyteStr, false, true},
		"empty":        {"", false, true},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, gotOK := nullbyte.Contains(tt.in)
			if tt.want != got || tt.wantOK != gotOK {
				t.Errorf("in=%s, want=%t, wantOK=%t, got=%t, gotOK=%t", tt.in, tt.want, tt.wantOK, got, gotOK)
			}
		})
	}
}

func TestContainsInByte(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		in []byte

		want   bool
		wantOK bool
	}{
		"null byte":    {[]byte(before), true, true},
		"no null byte": {[]byte(noNullbyteStr), false, true},
		"empty":        {[]byte(""), false, true},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, gotOK := nullbyte.Contains(tt.in)
			if tt.want != got || tt.wantOK != gotOK {
				t.Errorf("in=%s, want=%t, wantOK=%t, got=%t, gotOK=%t", tt.in, tt.want, tt.wantOK, got, gotOK)
			}
		})
	}
}

func TestContainsInOther(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		in any

		want   bool
		wantOK bool
	}{
		"bool":  {true, false, false},
		"int":   {1, false, false},
		"float": {0.1, false, false},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, gotOK := nullbyte.Contains(tt.in)
			if tt.want != got || tt.wantOK != gotOK {
				t.Errorf("in=%s, want=%t, wantOK=%t, got=%t, gotOK=%t", tt.in, tt.want, tt.wantOK, got, gotOK)
			}
		})
	}
}

func TestTrimNullByteFromStr(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		in string

		want   string
		wantOK bool
	}{
		"null byte":    {before, afterTrim, true},
		"no null byte": {noNullbyteStr, noNullbyteStr, true},
		"empty":        {"", "", true},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, gotOK := nullbyte.TrimNullByte(tt.in)
			if tt.want != got || tt.wantOK != gotOK {
				t.Errorf("in=%s, want=%s, wantOK=%t, got=%s, gotOK=%t", tt.in, tt.want, tt.wantOK, got, gotOK)
			}
		})
	}
}

func TestTrimNullByteFromBytes(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		in []byte

		want   []byte
		wantOK bool
	}{
		"null byte":    {[]byte(before), []byte(afterTrim), true},
		"no null byte": {[]byte(noNullbyteStr), []byte(noNullbyteStr), true},
		"empty":        {[]byte(""), []byte(""), true},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, gotOK := nullbyte.TrimNullByte(tt.in)
			if !cmp.Equal(tt.want, got) || tt.wantOK != gotOK {
				t.Errorf("in=%v, want=%v, wantOK=%t, got=%v, gotOK=%t", tt.in, tt.want, tt.wantOK, got, gotOK)
			}
		})
	}
}

func TestTrimNullByteInOther(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		in any

		want   any
		wantOK bool
	}{
		"bool":  {true, true, false},
		"int":   {1, 1, false},
		"float": {0.1, 0.1, false},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, gotOK := nullbyte.TrimNullByte(tt.in)
			if tt.want != got || tt.wantOK != gotOK {
				t.Errorf("in=%s, want=%t, wantOK=%t, got=%t, gotOK=%t", tt.in, tt.want, tt.wantOK, got, gotOK)
			}
		})
	}
}

func TestRemoveNullByteFromStr(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		in string

		want   string
		wantOK bool
	}{
		"null byte":    {before, noNullbyteStr, true},
		"no null byte": {noNullbyteStr, noNullbyteStr, true},
		"empty":        {"", "", true},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, gotOK := nullbyte.RemoveNullByte(tt.in)
			if tt.want != got || tt.wantOK != gotOK {
				t.Errorf("in=%s, want=%s, wantOK=%t, got=%s, gotOK=%t", tt.in, tt.want, tt.wantOK, got, gotOK)
			}
		})
	}
}

func TestRemoveNullByteFromBytes(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		in []byte

		want   []byte
		wantOK bool
	}{
		"null byte":    {[]byte(before), []byte(noNullbyteStr), true},
		"no null byte": {[]byte(noNullbyteStr), []byte(noNullbyteStr), true},
		"empty":        {[]byte(""), nil, true}, // TODO Confirm why want is nil
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, gotOK := nullbyte.RemoveNullByte(tt.in)
			if !cmp.Equal(tt.want, got) || tt.wantOK != gotOK {
				t.Errorf("in=%v, want=%v, wantOK=%t, got=%v, gotOK=%t", tt.in, tt.want, tt.wantOK, got, gotOK)
			}
		})
	}
}

func TestRemoveNullByteInOther(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		in any

		want   any
		wantOK bool
	}{
		"bool":  {true, true, false},
		"int":   {1, 1, false},
		"float": {0.1, 0.1, false},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, gotOK := nullbyte.RemoveNullByte(tt.in)
			if tt.want != got || tt.wantOK != gotOK {
				t.Errorf("in=%s, want=%t, wantOK=%t, got=%t, gotOK=%t", tt.in, tt.want, tt.wantOK, got, gotOK)
			}
		})
	}
}
