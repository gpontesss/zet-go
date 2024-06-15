package io

import "io"

// https://pkg.go.dev/os#File.Seek
const seekFromStart int = 0
const seekFromOffset int = 1

// ReadRange reads the specified range of the file and returns it in a buffer.
func ReadRange(rs io.ReadSeeker, from, to int64) ([]byte, error) {
	_, err := rs.Seek(from, seekFromStart)
	if err != nil {
		return nil, err
	}
	rbuf := make([]byte, to-from)
	length, err := rs.Read(rbuf)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return rbuf[0:length], nil
}
