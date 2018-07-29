package site

import "fmt"

type Guge struct {
	Alibaba
}

func (a *Guge) NewSite(url, seed string) *Guge {
	a.Seed = seed
	a.Url = fmt.Sprintf("https://www.google.com/search?q=%s&ie=UTF-8", a.Seed)

	return a
}
