package search

import (
	"strings"
	"testing"
)

func TestFindStrIter(t *testing.T) {
	bufSize := int64(1024)
	rs := strings.NewReader("aa bb aa cc dd aa")
	ls := NewLazySearcher(rs, bufSize)

	t.Run("Find multiple matches", func(t *testing.T) {
		search := "aa"
		expectedMatches := []int64{0, 6, 15}

		iter := FindStrIter(&ls, search)
		matches := []int64{}
		for iter.Next() {
			matches = append(matches, iter.Index())
			expectedMatch := expectedMatches[len(matches)-1]
			if expectedMatch != iter.Index() {
				t.Errorf("Match is not correct: %v != %v",
					expectedMatch, iter.Index())
			}
		}
		if len(matches) != len(expectedMatches) {
			t.Errorf("Not enough matches: %v != %v", matches, expectedMatches)
		}
	})
}
