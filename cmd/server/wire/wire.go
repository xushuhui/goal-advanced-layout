//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"

	"fm-suggest/internal/conf"
	"fm-suggest/internal/data"
	"fm-suggest/internal/handler"
	"fm-suggest/internal/server"
	"fm-suggest/internal/service"
	"fm-suggest/pkg/app"
	"fm-suggest/pkg/helper/sid"
	"fm-suggest/pkg/jwt"
	"fm-suggest/pkg/log"
	"fm-suggest/pkg/server/http"
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
)

var serviceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
)

var dataSet = wire.NewSet(
	data.NewDB,
	data.NewRedis,
	data.NewData,
	data.NewUserRepo,
)
var serverSet = wire.NewSet(
	server.NewHTTPServer,
	server.NewJob,
	server.NewTask,
)

// build App
func newApp(httpServer *http.Server, job *server.Job) *app.App {
	return app.NewApp(
		app.WithServer(httpServer, job),
		app.WithName("demo-server"),
	)
}

func NewWire(*conf.Server, *conf.Data, *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		dataSet,
		serviceSet,
		handlerSet,
		serverSet,
		sid.NewSid,
		jwt.NewJwt,
		newApp,
	))
}
