package zet

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gpontesss/zet-go/lib/search"
	yaml "gopkg.in/yaml.v3"
)

var bufSize int64 = 1024

// Zet docs here.
type Zet struct {
	file   *os.File
	kasten *Zettelkasten
}

// Meta docs here.
type Meta struct {
	ID   int64    `yaml:"id"`
	Tags []string `yaml:"tags"`
}

// NewZet docs here.
func NewZet(dirname string, kasten *Zettelkasten) (Zet, error) {
	filename := fmt.Sprintf("%s/README.md", dirname)
	file, err := os.Open(filename)
	return Zet{file: file, kasten: kasten}, err
}

// Content docs here.
func (zet *Zet) Content() (string, error) {
	var content []byte
	content, err := ioutil.ReadAll(zet.file)
	return string(content), err
}

// Metadata docs here.
func (zet *Zet) Metadata() (Meta, error) {
	ls := search.NewLazySearcher(zet.file, bufSize)
	ls.Reset()
	from := search.FindNextStr(&ls, "---")
	to := search.FindNextStr(&ls, "\n---\n")
	if from < 0 || to < 0 {
		return Meta{}, errors.New("Failed to read metadata: header is not properly formatted")
	}
	bs, err := ls.ReadRange(from+4, to)
	var meta Meta
	yaml.Unmarshal(bs, &meta)
	return meta, err
}
