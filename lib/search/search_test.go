package search

import (
	"strings"
	"testing"
)

func TestFindNextStr(t *testing.T) {
	bufSize := int64(1024)
	rs := strings.NewReader("aa bb cc dd ee")
	s := NewLazySearcher(rs, bufSize)
	s.Reset()

	var expected int64 = 3
	if idx := FindNextStr(&s, "bb"); idx != expected {
		t.Fatalf("Expected index %v, got %v", expected, idx)
	}
	expected = 9
	if idx := FindNextStr(&s, "dd"); idx != expected {
		t.Fatalf("Expected index %v, got %v", expected, idx)
	}

}
