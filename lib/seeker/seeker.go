package seeker

import (
	"io"
)

// https://pkg.go.dev/os#File.Seek
const seekFromStart int = 0
const seekFromOffset int = 1

type LazySearcher struct {
	rs             io.ReadSeeker
	page           int64
	buf            []byte
	bufSize, bufAt int64
}

func NewLazySearcher(file io.ReadSeeker, bufSize int64) LazySearcher {
	return LazySearcher{
		rs:      file,
		buf:     make([]byte, bufSize),
		bufSize: bufSize,
		page:    0}
}

func (fs *LazySearcher) isAtEnd() bool {
	return fs.bufAt+1 >= int64(len(fs.buf)) &&
		int64(len(fs.buf)) < fs.bufSize
}

func (fs *LazySearcher) offset() int64 { return (fs.page * fs.bufSize) + fs.bufAt }
func (fs *LazySearcher) current() byte { return fs.buf[fs.bufAt] }
func (fs *LazySearcher) advance() byte {
	fs.bufAt += 1
	if fs.bufAt == fs.bufSize {
		// TODO: deal with error
		fs.nextPage()
	}
	return fs.buf[fs.bufAt]
}

func (fs *LazySearcher) nextPage() error {
	fs.page += 1
	_, err := fs.rs.Seek(fs.bufSize, seekFromOffset)
	if err != nil {
		return err
	}
	n, err := fs.rs.Read(fs.buf)
	fs.buf = fs.buf[:n]
	fs.bufAt = 0
	return err
}

func (fs *LazySearcher) Reset() error {
	fs.page = 0
	_, err := fs.rs.Seek(0, seekFromStart)
	if err != nil {
		return err
	}
	n, err := fs.rs.Read(fs.buf)
	fs.buf = fs.buf[:n]
	fs.bufAt = 0
	return err
}

func (fs *LazySearcher) Find(b byte) int64 {
	fs.Reset()
	for !fs.isAtEnd() && b != fs.current() {
		fs.advance()
	}
	if fs.current() == b {
		return fs.offset()
	}
	return -1
}

func (fs *LazySearcher) FindStr(str string) int64 {
	fs.Reset()
	for !fs.isAtEnd() {
		matched := 0
		for matched < len(str) && fs.current() == str[matched] {
			matched += 1
			fs.advance()
		}
		if matched == len(str) {
			return fs.offset() - int64(len(str))
		}
		fs.advance()
	}
	return -1

}

func (fs *LazySearcher) ReadRange(from, to int64) ([]byte, error) {
	_, err := fs.rs.Seek(from, seekFromStart)
	if err != nil {
		return nil, err
	}
	rbuf := make([]byte, to-from)
	_, err = fs.rs.Read(rbuf)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return rbuf, nil
}
