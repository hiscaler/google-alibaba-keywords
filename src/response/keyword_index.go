package response

// keyword/index 接口数据结构
type KeywordIndex struct {
	Success bool `json:"success"`
	Data    Data `json:"data"`
}

type Data struct {
	Items []Items
}

type Items struct {
	SiteName string
	Id       int
	Name     string
}
