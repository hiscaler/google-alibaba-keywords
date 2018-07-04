package site

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"net/url"
	"crypto/tls"
	"time"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"log"
)

type Google struct {
	url          string
	seedKeyword  string
	Level        int
	Keywords     map[string]Keyword
	SearchResult map[string]Keyword
}

type SearchResult struct {
	id   int
	name string
}

func (g *Google) setUrl(url string) *Google {
	g.url = url

	return g
}

func (g *Google) getUrl() string {
	return g.url
}

func (g *Google) SetSeed(seedKeyword string) *Google {
	g.seedKeyword = seedKeyword
	url := fmt.Sprintf("https://www.google.com/search?q=%s&ie=UTF-8", seedKeyword)
	g.setUrl(url)

	return g
}

func (g *Google) Search() *Google {
	println(g.getUrl())
	request, _ := http.NewRequest("GET", g.getUrl(), nil)
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36")
	proxy, _ := url.Parse("http://127.0.0.1:1080")
	tr := &http.Transport{
		Proxy:           http.ProxyURL(proxy),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 5,
	}
	response, err := client.Do(request)
	if err == nil {
		if response.StatusCode == 200 {
			content, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Printf("Read response body error: %s", err)
			} else {
				doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(content)))
				if err == nil {
					doc.Find("#brs .card-section a").Each(func(i int, selection *goquery.Selection) {
						name := strings.Trim(selection.Text(), " ")
						url, _ := selection.Attr("href")
						fmt.Printf("#%02d Keyword name = '%s' url = %s\n", i+1, name, url)
						keyword := new(Keyword)
						if g.Level == 0 {
							keyword.Class = directoryKeyword
						} else {
							keyword.Class = adverbKeyword
						}
						keyword.Name = name
						if k, err := keyword.save(); err == nil {
							g.Keywords[name] = *k
						}
					})
					g.Level += 1
				} else {
					log.Fatal(err)
				}
			}
		}
		defer response.Body.Close()
	} else {
		fmt.Printf("HTTP error: %s", err)
	}

	return g
}
