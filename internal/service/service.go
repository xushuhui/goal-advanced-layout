package service

import (
	"github.com/google/wire"

	"goal-advanced-layout/pkg/helper/sid"
	"goal-advanced-layout/pkg/jwt"
	"goal-advanced-layout/pkg/log"
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
