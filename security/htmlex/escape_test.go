package htmlex_test

import (
	"testing"

	"github.com/exaream/go-lib/security/htmlex"
)

func TestEscapeString(t *testing.T) {
	t.Parallel()
	for in, want := range htmlex.ExportSpecialChars {
		got := htmlex.EscapeString(in)
		if want != got {
			t.Errorf("in: %s, want: %s, got: %s\n", in, want, got)
		}
	}
}

func TestUnescapeString(t *testing.T) {
	t.Parallel()
	for want, in := range htmlex.ExportSpecialChars {
		got := htmlex.UnescapeString(in)
		if want != got {
			t.Errorf("in: %s, want: %s, got: %s\n", in, want, got)
		}
	}
}
