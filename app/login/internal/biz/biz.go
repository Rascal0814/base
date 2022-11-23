package biz

import (
	userV1 "base/api/user/v1"
	"base/app/login/internal/conf"
	"context"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"time"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewLoginUsecase, NewUserServiceClient, NewRedisCmd, NewDiscovery)

func NewRedisCmd(conf *conf.Data, logger log.Logger) redis.Cmdable {
	log := log.NewHelper(log.With(logger, "module", "login-service/biz"))
	client := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		DialTimeout:  time.Second * 2,
		PoolSize:     10,
	})
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*10)
	defer cancelFunc()
	err := client.Ping(timeout).Err()
	if err != nil {
		log.Fatalf("redis connect error: %v", err)
	}
	return client
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewUserServiceClient(r registry.Discovery, tp *tracesdk.TracerProvider) userV1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///base.user"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(tracing.WithTracerProvider(tp)),
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := userV1.NewUserClient(conn)
	return c
}
