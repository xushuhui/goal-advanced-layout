package middleware

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"goal-advanced-layout/api"
	"goal-advanced-layout/pkg/jwt"
	"goal-advanced-layout/pkg/log"
)

func StrictAuth(j *jwt.JWT, logger *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("Authorization")
		if tokenString == "" {
			logger.WithContext(ctx).Warn("No token", slog.Any("data", map[string]any{
				"url":    ctx.Request.URL,
				"params": ctx.Params,
			}))
			api.Fail(ctx, http.StatusUnauthorized, api.ErrUnauthorized, nil)
			ctx.Abort()
			return
		}

		claims, err := j.ParseToken(tokenString)
		if err != nil {
			logger.WithContext(ctx).Error("token error", slog.Any("data", map[string]any{
				"url":    ctx.Request.URL,
				"params": ctx.Params,
				"err":    err,
			}))
			api.Fail(ctx, http.StatusUnauthorized, api.ErrUnauthorized, nil)
			ctx.Abort()
			return
		}

		ctx.Set("claims", claims)
		recoveryLoggerFunc(ctx, logger)
		ctx.Next()
	}
}

func NoStrictAuth(j *jwt.JWT, logger *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("Authorization")
		if tokenString == "" {
			tokenString, _ = ctx.Cookie("accessToken")
		}
		if tokenString == "" {
			tokenString = ctx.Query("accessToken")
		}
		if tokenString == "" {
			ctx.Next()
			return
		}

		claims, err := j.ParseToken(tokenString)
		if err != nil {
			ctx.Next()
			return
		}

		ctx.Set("claims", claims)
		recoveryLoggerFunc(ctx, logger)
		ctx.Next()
	}
}

func recoveryLoggerFunc(ctx *gin.Context, logger *log.Logger) {
	userInfo := ctx.MustGet("claims").(*jwt.MyCustomClaims)
	logger.WithValue(ctx, slog.String("UserId", userInfo.UserId))
}
