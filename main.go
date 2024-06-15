package main

import (
	"fmt"

	zet "github.com/gpontesss/zet-go/lib/zet"
)

func main() {
	dir := "/Users/guilherme.pontes1/zet"
	kasten, err := zet.NewZettelkasten(dir)
	if err != nil {
		fmt.Printf("Failed to create Zettelkasten: %s", err)
	}
	for _, zet := range kasten.Zets() {
		fmt.Println(zet.Meta.ID, zet.Title())
	}
}
