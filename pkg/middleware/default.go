package middleware

import (
	"base/pkg/jwt"
	"base/pkg/response"
	"context"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func NewWhiteListMatcher() selector.MatchFunc {

	whiteList := make(map[string]struct{})
	whiteList["/base.api.login.v1.Login/Login"] = struct{}{}
	whiteList["/base.api.login.v1.Register/Register"] = struct{}{}
	whiteList["/v1.Login/Login"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

var DefaultHttpMiddleWare = []http.ServerOption{
	http.Middleware(
		recovery.Recovery(),
		validate.Validator(),
		selector.Server(
			jwt.MiddleWare,
		).
			Match(NewWhiteListMatcher()).
			Build(),
	),
	http.ResponseEncoder(response.ResponseEncoder()),
	http.ErrorEncoder(response.ErrorEncode()),
}

var DefaultGrpcMiddleWare = []grpc.ServerOption{
	grpc.Middleware(
		recovery.Recovery(),
		validate.Validator(),
	),
}
