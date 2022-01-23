package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "mhdsg/api/server/service/v1"
	"mhdsg/app/server/internal/conf"
	"mhdsg/app/server/internal/service"
	"time"
)

// NewGRPCServer new a HTTP server.
func NewGRPCServer(c *conf.AppSettings, s *service.ServerService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			LogMiddleware(),
		),
	}
	//if c.Grpc.Network != "" {
	//	opts = append(opts, grpc.Network(c.Grpc.Network))
	//}
	if c.Server.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Server.Grpc.Addr))
	}
	if c.Server.Grpc.Timeout > 0 {
		opts = append(opts, grpc.Timeout(time.Duration(c.Server.Grpc.Timeout)*time.Second))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterServerServer(srv, s)
	return srv
}
