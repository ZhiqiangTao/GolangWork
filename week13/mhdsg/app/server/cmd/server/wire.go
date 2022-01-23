//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"mhdsg/app/server/internal/biz"
	"mhdsg/app/server/internal/conf"
	"mhdsg/app/server/internal/data"
	"mhdsg/app/server/internal/server"
	"mhdsg/app/server/internal/service"
)

// initApp init kratos application.
func initApp(*conf.AppSettings, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
