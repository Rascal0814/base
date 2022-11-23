package server

import (
	"base/api/user/v1"
	"base/app/user/internal/conf"
	"base/app/user/internal/service"
	"base/pkg/middleware"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, hs *service.UserService, logger log.Logger) *http.Server {
	var opts = middleware.DefaultHttpMiddleWare
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}

	srv := http.NewServer(opts...)
	v1.RegisterUserHTTPServer(srv, hs)
	return srv
}
