package server

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"nunu-http-layout/api"
	"nunu-http-layout/docs"
	"nunu-http-layout/internal/conf"
	"nunu-http-layout/internal/handler"
	"nunu-http-layout/internal/server/middleware"
	"nunu-http-layout/pkg/jwt"
	"nunu-http-layout/pkg/log"
	"nunu-http-layout/pkg/server/http"
)

func NewHTTPServer(
	logger *log.Logger,
	conf *conf.Server,
	jwt *jwt.JWT,
	userHandler handler.UserHandler,
) *http.Server {
	gin.SetMode(gin.DebugMode)
	s := http.NewServer(
		gin.Default(),
		logger,
		http.WithServerHost(conf.Http.Addr),
	)

	// swagger doc
	docs.SwaggerInfo.BasePath = "/api"
	s.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		ginSwagger.DefaultModelsExpandDepth(-1),
	))

	s.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
	)
	s.GET("/", func(ctx *gin.Context) {
		logger.WithContext(ctx).Info("hello")
		api.HandleSuccess(ctx, map[string]any{
			":)": "Thank you for using nunu!",
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
