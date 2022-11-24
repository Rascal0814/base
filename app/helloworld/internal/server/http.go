package server

import (
	v1 "base/api/helloworld/v1"
	"base/app/helloworld/internal/conf"
	"base/app/helloworld/internal/service"
	"base/pkg/middleware"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.HelloworldService, logger log.Logger) *http.Server {
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
	v1.RegisterHelloworldHTTPServer(srv, greeter)
	return srv
}
