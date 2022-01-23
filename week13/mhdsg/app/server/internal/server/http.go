package server

import (
	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "mhdsg/api/server/service/v1"
	"mhdsg/app/server/internal/conf"
	"mhdsg/app/server/internal/service"
	"mhdsg/pkg/res"
	http2 "net/http"
	"time"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.AppSettings, s *service.ServerService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			LogMiddleware(),
		),
	}
	//if c.Server.Http.Network != "" {
	//	opts = append(opts, http.Network(c.Http.Network))
	//}
	if c.Server.Http.Addr != "" {
		opts = append(opts, http.Address(c.Server.Http.Addr))
	}
	if c.Server.Http.Timeout > 0 {
		opts = append(opts, http.Timeout(time.Duration(c.Server.Http.Timeout)*time.Second))
	}
	opts = append(opts, http.ErrorEncoder(func(w http2.ResponseWriter, r *http2.Request, err error) {
		codec := encoding.GetCodec("json")
		raw := errors.FromError(err)
		e := res.Fail(raw.Code, raw.Reason, raw.Message)
		data, err := codec.Marshal(e)
		if err != nil {
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
		return
	}))

	srv := http.NewServer(opts...)
	v1.RegisterServerHTTPServer(srv, s)

	return srv
}
