package site

import (
	"fmt"
	"logger"
	"net/http"
	"net/url"
	"crypto/tls"
	"time"
	"io/ioutil"
)

// Alibaba
type Alibaba struct {
	url             string
	Seed            string
	QualifyKeywords []string // 限定词
	keywords        []string // 已有的关键词
	searchResult    []string // 搜索结果
	pageSource      []byte   // 页面源码
	Level           int
	Keywords        map[string]Keyword
}

func NewSite(url, seed string) *Alibaba {
	return &Alibaba{
		url:  fmt.Sprintf("https://www.alibaba.com/search?q=%s&ie=UTF-8", seed),
		Seed: seed,
	}
}

func (a *Alibaba) setUrl() *Alibaba {
	a.url = fmt.Sprintf("https://www.alibaba.com/search?q=%s&ie=UTF-8", a.Seed)

	return a
}

func (a *Alibaba) getUrl() string {
	return a.url
}

// Set `seed keyword` and search url
func (a *Alibaba) SetSeed(seed string) *Alibaba {
	a.Seed = seed
	a.setUrl()

	return a
}

func (a *Alibaba) SetQualifyKeywords(keywords []string) *Alibaba {
	a.QualifyKeywords = keywords

	return a
}

func (a *Alibaba) Search() *Alibaba {
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
				a.pageSource = content
			}
		}
		defer response.Body.Close()
	} else {
		logger.Instance.Error(fmt.Sprintf("HTTP error: %s", err))
	}

	return a
}

// 解析 Alibaba 爬取结果
func (a *Alibaba) Parse() *Alibaba {
	return a
}
