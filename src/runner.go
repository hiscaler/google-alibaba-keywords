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
	"logger"
)

func main() {
	keywordIndexResponse := response.KeywordIndex{}
	resp, err := http.Get(url.KeywordIndex())
	if err != nil {
		logger.Instance.Error(err.Error())
	} else {
		if resp.StatusCode == 200 {
			body, _ := ioutil.ReadAll(resp.Body)
			if err := json.Unmarshal([]byte(body), &keywordIndexResponse); err == nil {
				g := new(site.Google)
				g.Keywords = make(map[string]site.Keyword)
				for _, item := range keywordIndexResponse.Data.Items {
					logger.Instance.Info(fmt.Sprintf("Site Name = %s, Id = %d, Name = %s", item.SiteName, item.Id, item.Name))
					g.SetSeed(item.Name)
					g.Search()
					logger.Instance.Debug(fmt.Sprintf("%# v", pretty.Formatter(*g)))
				}
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Can't get keywords list.")
		}
	}

	defer logger.Instance.Close()
}
