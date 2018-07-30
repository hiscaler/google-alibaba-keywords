package site

import (
	"fmt"
	"logger"
	"net/http"
	"net/url"
	"crypto/tls"
	"time"
	"io/ioutil"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type ISite interface {
	NewSite(url, seed string)
	SetUrl()                              // 设置要爬取的页面地址
	SetSeed(seed string)                  // 设置种子词
	SetQualifyKeywords(keywords []string) // 设置限定词
	Search()                              // 爬取页面
	Parse()                               // 解析爬取的页面
	UpdateMeta()                          // 更新 page_keywords, page_description
	FetchProducts()                       // 根据关键词爬取商品数据
}

type Site struct {
	ISite
	url             string
	Seed            string
	QualifyKeywords []string // 限定词
	keywords        []string // 已有的关键词
	searchResult    []string // 搜索结果
	pageSource      string   // 页面源码
	Level           int
	Keywords        map[string]Keyword
}

//func NewSite() ISite {
//	return nil
//}

func (a *Site) getUrl() string {
	return a.url
}

//func (a *Site) SetUrl() {}

func (a *Site) GetUrl() string {
	return a.url
}

// Set `seed keyword` and search url
func (a *Site) SetSeed(seed string) {
	a.Seed = seed
}

func (a *Site) SetQualifyKeywords(keywords []string) {
	a.QualifyKeywords = keywords
}

func (a *Site) Search() *Site {
	logger.Instance.Info(a.getUrl())
	request, _ := http.NewRequest("GET", a.getUrl(), nil)
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
				a.pageSource = string(content)
			}
		}
		defer response.Body.Close()
	} else {
		logger.Instance.Error(fmt.Sprintf("HTTP error: %s", err))
	}

	return a
}

// Alibaba
type Ali struct {
	Site
}

func (a *Ali) SetUrl() {
	a.url = fmt.Sprintf("https://www.alibaba.com/search?q=%s&ie=UTF-8", a.Seed)
}

// Google
type Gg struct {
	Site
}

func (g *Gg) SetUrl() {
	g.url = fmt.Sprintf("https://www.google.com/search?q=%s&ie=UTF-8", g.Seed)
}

// 解析谷歌爬取结果
func (g *Gg) Parse() *Gg {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(g.pageSource)))
	if err == nil {
		doc.Find("#brs .card-section a").Each(func(i int, selection *goquery.Selection) {
			name := strings.Trim(selection.Text(), " ")
			// 检查是否在限定词中
			isValid := false
			for _, v := range g.QualifyKeywords {
				if v == name {
					isValid = true
					break
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

	return g
}
