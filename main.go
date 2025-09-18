package main

import (
	"context"
	"flag"
	"log/slog"

	"goal-advanced-layout/internal/server"
	"goal-advanced-layout/pkg/app"
	"goal-advanced-layout/pkg/config"
	"goal-advanced-layout/pkg/log"
	"goal-advanced-layout/pkg/server/http"
)

func newApp(httpServer *http.Server, job *server.Job) *app.App {
	return app.NewApp(
		app.WithServer(httpServer, job),
		app.WithName("server"),
	)
}

func main() {
	flagconf := flag.String("conf", "configs/dev.yaml", "config path, eg: -conf ./configs/dev.yaml")
	flag.Parse()
	conf := config.NewConfig(*flagconf)

	logger := log.NewLog()
	app, cleanup, err := NewWire(conf.Server, conf.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()
	logger.Info("server start", slog.String("host", conf.Server.HTTP.Addr))

	if err = app.Run(context.Background()); err != nil {
		panic(err)
	}
}
