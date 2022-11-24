// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"base/app/user/internal/biz"
	"base/app/user/internal/conf"
	"base/app/user/internal/data"
	"base/app/user/internal/server"
	"base/app/user/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	registrar := server.NewRegistrar(registry)
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	cmdable := data.NewRedisCmd(confData, logger)
	userRepo := data.NewGreeterRepo(dataData, logger, cmdable)
	user := biz.NewUse(userRepo, logger)
	userService := service.NewUserService(user)
	grpcServer := server.NewGRPCServer(confServer, userService, logger)
	httpServer := server.NewHTTPServer(confServer, userService, logger)
	app := newApp(logger, registrar, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
