package site

import "fmt"

// Alibaba
type Alibaba struct {
	Url             string
	Seed            string
	QualifyKeywords []string
	keywords        []string
	searchResult    []string
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
