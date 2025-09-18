package server

import (
	"goal-advanced-layout/api"
	"goal-advanced-layout/internal/conf"
	"goal-advanced-layout/internal/handler"
	"goal-advanced-layout/internal/server/middleware"
	"goal-advanced-layout/pkg/jwt"
	"goal-advanced-layout/pkg/log"
	"goal-advanced-layout/pkg/server/http"

	"github.com/gin-gonic/gin"
)

func NewHTTPServer(
	logger *log.Logger,
	conf *conf.Server,
	jwt *jwt.JWT,
	userHandler *handler.UserHandler,
) *http.Server {
	gin.SetMode(gin.DebugMode)
	s := http.NewServer(
		gin.Default(),
		logger,
		http.WithServerHost(conf.HTTP.Addr),
	)

	s.Use(
		middleware.CORS(),
		middleware.ResponseLog(logger),
		middleware.RequestLog(logger),
	)
	s.GET("/", func(ctx *gin.Context) {
		logger.WithContext(ctx).Info("hello")
		api.Succeed(ctx, map[string]any{
			":)": "Thank you!",
		})
	})

	api := s.Group("/api")
	{
		// No route group has permission
		noAuthRouter := api.Group("/")
		{
			noAuthRouter.POST("/register", userHandler.Register)
			noAuthRouter.POST("/login", userHandler.Login)
		}
		// Non-strict permission routing group
		noStrictAuthRouter := api.Group("/").Use(middleware.NoStrictAuth(jwt, logger))
		{
			noStrictAuthRouter.GET("/user", userHandler.GetProfile)
		}

		// Strict permission routing group
		strictAuthRouter := api.Group("/").Use(middleware.StrictAuth(jwt, logger))
		{
			strictAuthRouter.PUT("/user", userHandler.UpdateProfile)
		}
	}

	return s
}
