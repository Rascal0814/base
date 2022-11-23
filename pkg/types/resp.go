package types

// Resp is the http response struct
type Resp struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	PageCtx *PageCtx    `json:"page_ctx,omitempty"`
}
