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
	zet := kasten.Zets()[0]
	content, err := zet.Content()
	if err != nil {
		fmt.Printf("Failed to read zet content: %s", err)
	}
	fmt.Println(content)
	meta, err := zet.Metadata()
	if err != nil {
		fmt.Printf("Failed to get zet metadata: %s", err)
	}
	fmt.Println(meta)
}
