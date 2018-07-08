package main

import (
	"fmt"
	"net/http"
	"url"
	"errors"
	"io/ioutil"
	"encoding/json"
	"site"
	"response"
)

func main() {
	keywordIndexResponse := response.KeywordIndex{}
	resp, err := http.Get(url.KeywordIndex())
	if err != nil {
		errors.New(err.Error())
	}
	if resp.StatusCode == 200 {
		fmt.Println(resp.Body)
		body, _ := ioutil.ReadAll(resp.Body)
		if err := json.Unmarshal([]byte(body), &keywordIndexResponse); err == nil {
			g := new(site.Google)
			g.Keywords = make(map[string]site.Keyword)
			for _, item := range keywordIndexResponse.Data.Items {
				fmt.Printf("Site Name = %s, Id = %d, Name = %s", item.SiteName, item.Id, item.Name)
				g.SetSeed(item.Name)
				g.Search()
				fmt.Println(*g)
			}
		} else {
			fmt.Println(err)
		}
	} else {
		errors.New("Can't get keywords list.")
	}
}
