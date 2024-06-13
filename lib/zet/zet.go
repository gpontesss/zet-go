package zet

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Zet struct {
	filename string
	file     *os.File
	metadata map[string]string
	kasten   *Zettelkasten
}

type Zettelkasten struct {
	dir  string
	zets []Zet
}

func (zet *Zet) Content() (string, error) {
	var content []byte
	content, err := ioutil.ReadAll(zet.file)
	return string(content), err
}

func (zet *Zet) Metadata() (string, error) {
	return "", nil
}

func NewZet(dirname string, kasten *Zettelkasten) (Zet, error) {
	filename := fmt.Sprintf("%s/README.md", dirname)
	file, err := os.Open(filename)
	return Zet{filename: filename, file: file, kasten: kasten}, err
}

func NewZettelkasten(dirname string) (*Zettelkasten, error) {
	dirs, err := os.ReadDir(dirname)
	if err != nil {
		return nil, err
	}
	kasten := &Zettelkasten{
		dir:  dirname,
		zets: make([]Zet, 0, len(dirs))}
	for _, dir := range dirs {
		// TODO: properly load only dirs with expected timestamp name
		if dir.Name() == ".git" || dir.Name() == "buffer" {
			continue
		}
		zet, err := NewZet(fmt.Sprintf("%s/%s", kasten.dir, dir.Name()), kasten)
		kasten.zets = append(kasten.zets, zet)
		if err != nil {
			return nil, fmt.Errorf("Failed to read zet %s: %s", dir.Name(), err)
		}
	}
	return kasten, nil
}

func (kasten *Zettelkasten) Zets() []Zet { return kasten.zets }
