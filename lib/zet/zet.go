package zet

import (
	"errors"
	"fmt"
	"os"

	_io "github.com/gpontesss/zet-go/lib/io"
	"github.com/gpontesss/zet-go/lib/md"
	"github.com/gpontesss/zet-go/lib/search"
	"github.com/yuin/goldmark"
	mdast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
	yaml "gopkg.in/yaml.v3"
)

var bufSize int64 = 1024

// Meta docs here.
type Meta struct {
	ID   int64    `yaml:"id"`
	Tags []string `yaml:"tags"`
}

// Zet docs here.
type Zet struct {
	file      *os.File
	kasten    *Zettelkasten
	mdTree    mdast.Node
	mdContent []byte
	Meta      Meta
}

// NewZet docs here.
func NewZet(dirname string, kasten *Zettelkasten) (Zet, error) {
	filename := fmt.Sprintf("%s/README.md", dirname)
	file, err := os.Open(filename)
	if err != nil {
		return Zet{}, fmt.Errorf("Failed to read zet %v file: %w", filename, err)
	}
	zet := Zet{file: file, kasten: kasten}
	if err := zet.parseContent(); err != nil {
		return Zet{}, fmt.Errorf("Failed to parse zet %v content: %w", filename, err)
	}
	return zet, nil
}

// Title docs here.
func (zet *Zet) Title() string {
	header := md.FirstHeading(zet.mdTree)
	if header != nil {
		return string(header.Text(zet.mdContent))
	}
	return "<titleless>"
}

func (zet *Zet) parseContent() error {
	ls := search.NewLazySearcher(zet.file, bufSize)
	ls.Reset()
	metaStart := search.FindNextStr(&ls, "---") + 4
	metaEnd := search.FindNextStr(&ls, "\n---\n")
	if metaStart < 0 || metaEnd < 0 {
		return errors.New("Can't to delimit header metadata")
	}
	if err := zet.parseMeta(metaStart, metaEnd); err != nil {
		return fmt.Errorf("Faile to parse metadata: %w", err)
	}
	mdStart := metaEnd + 4
	if err := zet.parseMd(mdStart); err != nil {
		return fmt.Errorf("Failed to parse markdown content: %w", err)
	}
	return nil
}

func (zet *Zet) parseMd(startOffset int64) error {
	var err error
	zet.mdContent, err = _io.ReadRange(
		zet.file, startOffset, search.LazySeqLen(zet.file, bufSize))
	if err != nil {
		return err
	}
	zet.mdTree = goldmark.New().Parser().Parse(text.NewReader(zet.mdContent))
	return nil
}

func (zet *Zet) parseMeta(from, to int64) error {
	bs, err := _io.ReadRange(zet.file, from, to)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(bs, &zet.Meta)
}
