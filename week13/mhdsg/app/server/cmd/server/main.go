package main

import (
	"flag"
	"github.com/go-kratos/kratos/v2/encoding/json"
	"os"

	"mhdsg/app/server/internal/conf"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "D:\\上海元聚\\mhd.sg.micro\\mhdsg\\app\\server\\configs", "config path, eg: -conf config.yaml")
}

//D:\上海元聚\mhd.sg.micro\mhdsg\app\server\cmd\server\main.go
//D:\上海元聚\mhd.sg.micro\mhdsg\app\server\configs\appsettings.json

func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
			gs,
		),
	)
}

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var config conf.AppSettings
	if err := c.Scan(&config); err != nil {
		panic(err)
	}
	app, cleanup, err := initApp(&config, logger)
	//	app, cleanup, err := initApp(bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	json.MarshalOptions.UseProtoNames = true
	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
