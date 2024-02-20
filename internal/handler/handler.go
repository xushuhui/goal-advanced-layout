package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"goal-advanced-layout/pkg/jwt"
	"goal-advanced-layout/pkg/log"
)

var ProviderSet = wire.NewSet(NewHandler, NewUserHandler)

type Handler struct {
	logger *log.Logger
}

func NewHandler(logger *log.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}

func GetUserIdFromCtx(ctx *gin.Context) string {
	v, exists := ctx.Get("claims")
	if !exists {
		return ""
	}
	return v.(*jwt.MyCustomClaims).UserId
}
