package seeker

import (
	"strings"
	"testing"
)

func TestLazySearcher(t *testing.T) {
	bufSize := int64(1024)
	rs := strings.NewReader("abcdef")
	fs := NewLazySearcher(rs, bufSize)

	t.Run("Test Find", func(t *testing.T) {
		t.Run("Find byte that is in the string", func(t *testing.T) {
			if 0 > fs.Find('f') {
				t.Fatalf("Expected 'f' to be in string")
			}
		})

		t.Run("Find byte that is NOT in the string", func(t *testing.T) {
			if 0 <= fs.Find('g') {
				t.Fatalf("Expected 'g' to NOT be in string")
			}
		})
	})

	t.Run("Test FindStr", func(t *testing.T) {
		t.Run("Find substring", func(t *testing.T) {
			if 0 > fs.FindStr("abc") {
				t.Fatalf("Expected 'abc' to be a substring")
			}
		})
		t.Run("Do NOT find substring", func(t *testing.T) {
			if 0 <= fs.FindStr("cba") {
				t.Fatalf("Didn't expected to find 'cba' a a substring")
			}
		})
	})
}
