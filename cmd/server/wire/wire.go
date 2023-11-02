//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"

	"nunu-http-layout/internal/conf"
	"nunu-http-layout/internal/data"
	"nunu-http-layout/internal/handler"
	"nunu-http-layout/internal/server"
	"nunu-http-layout/internal/service"
	"nunu-http-layout/pkg/app"
	"nunu-http-layout/pkg/helper/sid"
	"nunu-http-layout/pkg/jwt"
	"nunu-http-layout/pkg/log"
	"nunu-http-layout/pkg/server/http"
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
		data.ProviderSet,
		service.ProviderSet,
		handler.ProviderSet,
		server.ProviderSet,
		sid.NewSid,
		jwt.NewJwt,
		newApp,
	))
}
