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
	"logger"
)

type Google struct {
	url             string
	seedKeyword     string
	Level           int
	Keywords        map[string]Keyword
	SearchResult    map[string]Keyword
	QualifyKeywords []string
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

// 设置限定词
func (g *Google) SetQualifyKeywords(keyowrds []string) *Google {
	g.QualifyKeywords = keyowrds

	return g
}

func (g *Google) Search() *Google {
	logger.Instance.Info(g.getUrl())
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
				logger.Instance.Debug(fmt.Sprintf("Read response body error: %s", err))
			} else {
				doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(content)))
				if err == nil {
					doc.Find("#brs .card-section a").Each(func(i int, selection *goquery.Selection) {
						name := strings.Trim(selection.Text(), " ")
						// 检查是否在限定词中
						isValid := false
						for _, v := range g.QualifyKeywords {
							if v == name {
								isValid = true
								break;
							}
						}
						if isValid {
							url, _ := selection.Attr("href")
							logger.Instance.Info(fmt.Sprintf("#%02d Keyword name = '%s' url = %s", i+1, name, url))
							keyword := new(Keyword)
							if g.Level == 0 {
								keyword.Class = directoryKeyword
							} else {
								keyword.Class = adverbKeyword
							}
							keyword.Name = name
							if k, err := keyword.Save(); err == nil {
								g.Keywords[name] = *k
							}
						}
					})
					g.Level += 1
				} else {
					logger.Instance.Info(err.Error())
				}
			}
		}
		defer response.Body.Close()
	} else {
		logger.Instance.Error(fmt.Sprintf("HTTP error: %s", err))
	}

	return g
}
