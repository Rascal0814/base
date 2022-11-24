package types

// Resp is the http response struct
type Resp struct {
	Code    int         `json:"code"` //200 success
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	PageCtx *PageCtx    `json:"page_ctx,omitempty"`
}
