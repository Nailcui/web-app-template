package web

type Page struct {
	PageNo   int32       `json:"pageNo"`
	PageSize int32       `json:"pageSize"`
	Total    int32       `json:"total"`
	List     interface{} `json:"list"`
}
