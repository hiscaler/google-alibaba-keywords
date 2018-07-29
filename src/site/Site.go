package site

type Site interface {
	NewSite(url, seed string)
	SetUrl(url string)
	SetSeed(seed string)
	SetQualifyKeywords(keywords []string)
	Search()
}
