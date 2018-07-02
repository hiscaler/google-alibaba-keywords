package main

import (
	"site"
	"fmt"
)

func main() {
	seedKeywords := []string{"Bags"}
	g := new(site.Google)
	for _, seedKeyword := range seedKeywords {
		g.SetSeed(seedKeyword)
		g.Search()
		fmt.Println(*g)
	}
}
