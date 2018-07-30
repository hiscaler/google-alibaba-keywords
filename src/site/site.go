package site

type Site interface {
	NewSite(url, seed string)
	SetUrl(url string)                    // 设置要爬取的页面地址
	SetSeed(seed string)                  // 设置种子词
	SetQualifyKeywords(keywords []string) // 设置限定词
	Search()                              // 爬取页面
	Parse()                               // 解析爬取的页面
	UpdateMeta()                          // 更新 page_keywords, page_description
	FetchProducts()                       // 根据关键词爬取商品数据
}
