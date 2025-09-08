package biz

import (
	"github.com/google/wire"

	"goal-advanced-layout/pkg/helper/sid"
	"goal-advanced-layout/pkg/jwt"
	"goal-advanced-layout/pkg/log"
)

var ProviderSet = wire.NewSet(NewUsecase, NewUserUsecase)

type Usecase struct {
	logger *log.Logger
	sid    *sid.Sid
	jwt    *jwt.JWT
}

func NewUsecase(logger *log.Logger, sid *sid.Sid, jwt *jwt.JWT) *Usecase {
	return &Usecase{
		logger: logger,
		sid:    sid,
		jwt:    jwt,
	}
}
