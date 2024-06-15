package search

import (
	"io"
)

// https://pkg.go.dev/os#File.Seek
const seekFromStart int = 0
const seekFromOffset int = 1

// lazySearcher implements the Searcher interface and complementary high-level
// search methods.
type lazySearcher struct {
	rs             io.ReadSeeker
	page           int64
	buf            []byte
	bufSize, bufAt int64
}

// NewLazySearcher creates a new LazySearcher that will have a buffer with
// length = bufSize and will perform lazy searches on rs.
func NewLazySearcher(rs io.ReadSeeker, bufSize int64) lazySearcher {
	return lazySearcher{
		rs:      rs,
		buf:     make([]byte, bufSize),
		bufSize: bufSize,
		page:    0}
}

// Implements the Searcher interface
func (ls *lazySearcher) Offset() int64 {
	return (ls.page * ls.bufSize) + ls.bufAt
}

func (ls *lazySearcher) Current() byte {
	return ls.buf[ls.bufAt]
}

func (ls *lazySearcher) AtEnd() bool {
	return ls.bufAt+1 >= int64(len(ls.buf)) &&
		int64(len(ls.buf)) < ls.bufSize
}

func (ls *lazySearcher) Advance() bool {
	ls.bufAt += 1
	if ls.bufAt == ls.bufSize {
		// TODO: deal with error
		ls.nextPage()
	}
	return ls.bufAt < int64(len(ls.buf))
}

// TODO: how to deal with errors?
func (ls *lazySearcher) Reset() {
	ls.page = 0
	_, err := ls.rs.Seek(0, seekFromStart)
	if err != nil {
		return
	}
	n, err := ls.rs.Read(ls.buf)
	ls.buf = ls.buf[:n]
	ls.bufAt = 0
	return
}

func (ls *lazySearcher) nextPage() error {
	ls.page += 1
	_, err := ls.rs.Seek(ls.bufSize, seekFromOffset)
	if err != nil {
		return err
	}
	n, err := ls.rs.Read(ls.buf)
	ls.buf = ls.buf[:n]
	ls.bufAt = 0
	return err
}

// Find lazily searches for a byte in the file. It returns the first match. If
// there's no match, it returns -1.
func (ls *lazySearcher) Find(b byte) int64 {
	ls.Reset()
	for !ls.AtEnd() && b != ls.Current() {
		ls.Advance()
	}
	if ls.Current() == b {
		return ls.Offset()
	}
	return -1
}

// Find lazily searches for a string in the file. It returns the first match. If
// there's no match, it returns -1.
func (ls *lazySearcher) FindStr(str string) int64 {
	ls.Reset()
	iter := FindStrIter(ls, str)
	if !iter.Next() {
		return -1
	}
	return iter.Index()
}

// LazySeqLen docs here.
func LazySeqLen(rs io.ReadSeeker, bufSize int64) int64 {
	ls := NewLazySearcher(rs, bufSize)
	return SeqLen(&ls)
}
