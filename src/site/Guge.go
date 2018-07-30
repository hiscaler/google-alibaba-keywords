package site

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"logger"
)

type Guge struct {
	Alibaba
}

// 设置搜索链接地址
func (g *Guge) setUrl() *Guge {
	g.url = fmt.Sprintf("https://www.google.com/search?q=%s&ie=UTF-8", g.Seed)

	return g
}

// 解析谷歌爬取结果
func (g *Guge) Parse() *Guge {
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
