package service

import (
	"github.com/google/wire"

	"nunu-http-layout/pkg/helper/sid"
	"nunu-http-layout/pkg/jwt"
	"nunu-http-layout/pkg/log"
)

var ProviderSet = wire.NewSet(NewService,NewUserService)

type Service struct {
	logger *log.Logger
	sid    *sid.Sid
	jwt    *jwt.JWT
}

func NewService(logger *log.Logger, sid *sid.Sid, jwt *jwt.JWT) *Service {
	return &Service{
		logger: logger,
		sid:    sid,
		jwt:    jwt,
	}
}
