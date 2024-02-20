package main

import (
	"context"
	"flag"
	"fmt"

	"go.uber.org/zap"
	"goal-advanced-layout/internal/server"
	"goal-advanced-layout/pkg/app"
	"goal-advanced-layout/pkg/config"
	"goal-advanced-layout/pkg/log"
	"goal-advanced-layout/pkg/server/http"
)

func newApp(httpServer *http.Server, job *server.Job) *app.App {
	return app.NewApp(
		app.WithServer(httpServer, job),
		app.WithName("demo-server"),
	)
}

// @title           Nunu Example API
// @version         1.0.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8000
// @securityDefinitions.apiKey Bearer
// @in header
// @name Authorization
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
	logger.Info("server start", zap.String("host", conf.Server.Http.Addr))
	logger.Info("docs addr", zap.String("addr", fmt.Sprintf("%s/swagger/index.html", conf.Server.Http.Addr)))

	if err = app.Run(context.Background()); err != nil {
		panic(err)
	}
}
