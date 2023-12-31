// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
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

// Injectors from wire.go:

func NewWire(confServer *conf.Server, confData *conf.Data, logger *log.Logger) (*app.App, func(), error) {
	jwtJWT := jwt.NewJwt(confServer)
	handlerHandler := handler.NewHandler(logger)
	sidSid := sid.NewSid()
	serviceService := service.NewService(logger, sidSid, jwtJWT)
	db := data.NewDB(confData, logger)
	client := data.NewRedis(confData)
	dataData := data.NewData(db, client, logger)
	userRepo := data.NewUserRepo(dataData)
	userService := service.NewUserService(serviceService, userRepo)
	userHandler := handler.NewUserHandler(handlerHandler, userService)
	httpServer := server.NewHTTPServer(logger, confServer, jwtJWT, userHandler)
	job := server.NewJob(logger)
	appApp := newApp(httpServer, job)
	return appApp, func() {
	}, nil
}

// wire.go:

// build App
func newApp(httpServer *http.Server, job *server.Job) *app.App {
	return app.NewApp(app.WithServer(httpServer, job), app.WithName("demo-server"))
}
