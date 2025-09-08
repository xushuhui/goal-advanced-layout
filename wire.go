//go:build wireinject
// +build wireinject

package main

import (
	"goal-advanced-layout/internal/conf"
	"goal-advanced-layout/internal/data"
	"goal-advanced-layout/internal/handler"
	"goal-advanced-layout/internal/server"
	"goal-advanced-layout/internal/service"
	"goal-advanced-layout/pkg/app"
	"goal-advanced-layout/pkg/helper/sid"
	"goal-advanced-layout/pkg/jwt"
	"goal-advanced-layout/pkg/log"
	"goal-advanced-layout/pkg/server/http"

	"github.com/google/wire"
)

// build App

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
