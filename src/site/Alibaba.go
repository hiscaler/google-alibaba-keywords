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
	Url             string
	Seed            string
	QualifyKeywords []string
	keywords        []string
	searchResult    []string
	pageSource      []byte
	Level           int
	Keywords        map[string]Keyword
}

func (a *Alibaba) NewSite(url, seed string) *Alibaba {
	a.Seed = seed
	a.Url = fmt.Sprintf("https://www.alibaba.com/search?q=%s&ie=UTF-8", a.Seed)

	return a
}

func (a *Alibaba) SetUrl(url string) *Alibaba {
	a.Url = url

	return a
}

func (a *Alibaba) SetSeed(seed string) *Alibaba {
	a.Seed = seed

	return a
}

func (a *Alibaba) SetQualifyKeywords(keywords []string) *Alibaba {
	a.QualifyKeywords = keywords

	return a
}

func (a *Alibaba) Search() *Alibaba {
	logger.Instance.Info(a.Url)
	request, _ := http.NewRequest("GET", a.Url, nil)
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

func (a *Alibaba) Parse(content string) *Alibaba {
	if len(content) > 0 {

	}

	return a
}
