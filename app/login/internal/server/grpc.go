package server

import (
	v1 "base/api/login/v1"
	"base/app/login/internal/conf"
	"base/app/login/internal/service"
	"base/pkg/middleware"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, lg *service.LoginService, logger log.Logger) *grpc.Server {
	var opts = middleware.DefaultGrpcMiddleWare
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterLoginServer(srv, lg)
	return srv
}
