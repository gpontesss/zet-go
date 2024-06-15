package zet

import (
	"errors"
	"fmt"
	"os"

	"github.com/gpontesss/zet-go/lib/search"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
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

func (zet *Zet) Title() string {
	content, tree := zet.markdownTree()
	var header ast.Node
	ast.Walk(tree, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if n.Kind().String() == "Heading" {
			header = n
			return ast.WalkStop, nil
		}
		return ast.WalkContinue, nil
	})
	if header != nil {
		return string(header.Text(content))
	}
	return "<titleless>"
}

func (zet *Zet) markdownTree() ([]byte, ast.Node) {
	md := goldmark.New()
	ls := search.NewLazySearcher(zet.file, bufSize)
	ls.Reset()
	search.FindNextStr(&ls, "---")
	from := search.FindNextStr(&ls, "\n---\n")
	for ls.Advance() {
	}
	var content []byte
	content, err := ls.ReadRange(from+4, ls.Offset())
	if err != nil {
		panic(err)
	}
	return content, md.Parser().Parse(text.NewReader(content))
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
