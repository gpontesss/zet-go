package zet

import (
	"fmt"
	"os"
)

// Zettelkasten docs here.
type Zettelkasten struct {
	dir  string
	zets []Zet
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
