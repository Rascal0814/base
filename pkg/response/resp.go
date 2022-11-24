package response

import (
	"base/pkg/types"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/errors"
	Khttp "github.com/go-kratos/kratos/v2/transport/http"
	"net/http"
	"strings"
)

// Error 通常错误数据处理
func Error(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, types.Resp{
		Code:    code,
		Message: msg,
	})
}

// OK 通常成功数据处理
func OK(c *gin.Context, data interface{}, msg string) {
	c.JSON(http.StatusOK, types.Resp{
		Data:    data,
		Message: msg,
		Code:    http.StatusOK,
	})
}

// PageOK 分页数据处理
func PageOK(c *gin.Context, result interface{}, count int, pageIndex int, pageSize int, msg string) {
	c.JSON(http.StatusOK, types.Resp{
		Code:    http.StatusOK,
		Message: msg,
		Data:    result,
		PageCtx: &types.PageCtx{
			Index: pageIndex,
			Size:  pageSize,
			Total: count,
		},
	})
}

// ResponseEncoder 统一数据处理
func ResponseEncoder() Khttp.EncodeResponseFunc {
	return func(w http.ResponseWriter, r *http.Request, v interface{}) error {
		if v == nil {
			return nil
		}
		if rd, ok := v.(Khttp.Redirector); ok {
			url, code := rd.Redirect()
			http.Redirect(w, r, url, code)
			return nil
		}
		codec, _ := Khttp.CodecForRequest(r, "Accept")
		res := types.Resp{
			Code:    http.StatusOK,
			Message: "success",
			Data:    v,
		}
		data, err := codec.Marshal(res)
		if err != nil {
			return err
		}
		w.Header().Set("Content-Type", strings.Join([]string{"application", codec.Name()}, "/"))
		_, err = w.Write(data)
		if err != nil {
			return err
		}
		return nil
	}
}

// ErrorEncode 统一错误处理
func ErrorEncode() Khttp.EncodeErrorFunc {
	return func(w http.ResponseWriter, r *http.Request, err error) {
		se := errors.FromError(err)
		codec, _ := Khttp.CodecForRequest(r, "Accept")
		res := types.Resp{
			Code:    int(se.Code),
			Message: se.Message,
		}
		body, err := codec.Marshal(res)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", strings.Join([]string{"application", codec.Name()}, "/"))
		w.WriteHeader(int(se.Code))
		_, _ = w.Write(body)
	}
}
