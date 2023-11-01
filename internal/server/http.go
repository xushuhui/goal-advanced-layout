package server

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"fm-suggest/api"
	"fm-suggest/docs"
	"fm-suggest/internal/conf"
	"fm-suggest/internal/handler"
	"fm-suggest/internal/server/middleware"
	"fm-suggest/pkg/jwt"
	"fm-suggest/pkg/log"
	"fm-suggest/pkg/server/http"
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
		// ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", conf.GetInt("app.http.port"))),
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
