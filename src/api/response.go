package api

type KeywordIndexResponse struct {
	Success bool `json:"success"`
	Data    keywordIndexData `json:"data"`
}

type keywordIndexData struct {
	Items []KeywordItems
}

type KeywordItems struct {
	SiteName string
	Id       int
	Name     string
}
