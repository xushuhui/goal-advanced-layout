package main

import (
	"context"
	"flag"
	"fmt"

	"go.uber.org/zap"

	"nunu-http-layout/cmd/server/wire"
	"nunu-http-layout/pkg/config"
	"nunu-http-layout/pkg/log"
)

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
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	flagconf := flag.String("conf", "configs/dev.yml", "config path, eg: -conf ./configs/dev.yml")
	flag.Parse()

	conf := config.NewConfig(*flagconf)

	logger := log.NewLog()

	app, cleanup, err := wire.NewWire(conf.Server, conf.Data, logger)
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
