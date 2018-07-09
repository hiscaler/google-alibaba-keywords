package main

import (
	"fmt"
	"net/http"
	"url"
	"io/ioutil"
	"encoding/json"
	"site"
	"response"
	"github.com/kr/pretty"
)

func main() {
	keywordIndexResponse := response.KeywordIndex{}
	resp, err := http.Get(url.KeywordIndex())
	if err != nil {
		fmt.Println(err)
	} else {
		if resp.StatusCode == 200 {
			fmt.Println(resp.Body)
			body, _ := ioutil.ReadAll(resp.Body)
			if err := json.Unmarshal([]byte(body), &keywordIndexResponse); err == nil {
				g := new(site.Google)
				g.Keywords = make(map[string]site.Keyword)
				for _, item := range keywordIndexResponse.Data.Items {
					fmt.Printf("Site Name = %s, Id = %d, Name = %s\r\n", item.SiteName, item.Id, item.Name)
					g.SetSeed(item.Name)
					g.Search()
					fmt.Println(pretty.Formatter(*g))
				}
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Can't get keywords list.")
		}
	}

}
