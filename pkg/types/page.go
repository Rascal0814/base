package types

type Page struct {
	Index int `json:"index"`
	Size  int `json:"size"`
}

type PageCtx struct {
	Index int `json:"index"`
	Size  int `json:"size"`
	Total int `json:"total"`
}
