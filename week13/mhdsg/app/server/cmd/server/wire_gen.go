// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"mhdsg/app/server/internal/biz"
	"mhdsg/app/server/internal/conf"
	"mhdsg/app/server/internal/data"
	"mhdsg/app/server/internal/server"
	"mhdsg/app/server/internal/service"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(appSettings *conf.AppSettings, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(appSettings, logger)
	if err != nil {
		return nil, nil, err
	}
	iServerRepository := data.NewServerRepository(dataData, logger)
	serverUseCase := biz.NewServerUseCase(iServerRepository, logger, appSettings)
	serverService := service.NewServerService(serverUseCase)
	httpServer := server.NewHTTPServer(appSettings, serverService, logger)
	grpcServer := server.NewGRPCServer(appSettings, serverService, logger)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
