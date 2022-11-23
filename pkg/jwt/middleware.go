package jwt

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"strings"
)

// MiddleWare √
func MiddleWare(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		if header, ok := transport.FromServerContext(ctx); ok {
			auths := strings.SplitN(header.RequestHeader().Get(AuthorizationKey), " ", 2)
			if len(auths) != 2 || !strings.EqualFold(auths[0], BearerWord) {
				return nil, ErrMissingJwtToken
			}
			jwtToken := auths[1]
			_, _, err := JwtDecrypt(jwtToken)
			if err != nil {
				return nil, err
			}
			return handler(ctx, req)
		}
		return nil, ErrWrongContext
	}
}
