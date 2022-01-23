package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer)

func LogMiddleware() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				// Do something on entering
				defer func() {
					// Do something on exiting
				}()

				tr.RequestHeader().Set("taozhiqiang", "aaaaaaaaaa")
			}

			r, err := handler(ctx, req)
			return r, err
		}
	}
}
