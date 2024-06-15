package search

type StateSearcher interface {
	// Reset resets the sequence index to its beginning.
	Reset()
	// Current returns the current byte the sequence is located at.
	Current() byte
	// Offset returns the current offset of sequence being searched.
	Offset() int64
	// AtEnd tells if the sequence has been exausted.
	AtEnd() bool
	// Advance tries to advance the index of the sequence and return true if it
	// was to to advance; false if it wasn't.
	Advance() bool
}

// FindNextStr docs here.
func FindNextStr(s StateSearcher, str string) int64 {
	for !s.AtEnd() {
		matched := 0
		for matched < len(str) && s.Current() == str[matched] {
			matched += 1
			if !s.Advance() {
				break
			}
		}
		if matched == len(str) {
			return s.Offset() - int64(len(str))
		}
		s.Advance()
	}
	return -1
}

// Returns the length of the sequence attached to a StateSearcher.
func SeqLen(s StateSearcher) int64 {
	s.Reset()
	for s.Advance() {
	}
	return s.Offset()
}
