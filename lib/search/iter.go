package search

// IndexIter defines a interface for lazily searching indexes in a buffer, file,
// string etc. Next should return an index, or -1 for no further match.
type IndexIter interface {
	Next() bool
	Index() int64
}

type findStrIter struct {
	StateSearcher
	str string
}

// FindStrIter docs here.
func FindStrIter(searcher StateSearcher, str string) IndexIter {
	searcher.Reset()
	return findStrIter{searcher, str}
}

// Next docs here.
func (iter findStrIter) Next() bool {
	return FindNextStr(iter, iter.str) >= 0
}

// Index docs here.
func (iter findStrIter) Index() int64 {
	return iter.Offset() - int64(len(iter.str))
}
