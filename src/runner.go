package main

import (
	"site"
	"fmt"
)

func main() {
	seedKeywords := []string{"Bags"}
	g := new(site.Google)
	g.Keywords = make(map[string]site.Keyword)
	for _, seedKeyword := range seedKeywords {
		g.SetSeed(seedKeyword)
		g.Search()
		fmt.Println(*g)
	}
}
