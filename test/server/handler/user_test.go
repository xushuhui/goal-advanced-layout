package handler

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	v1 "goal-advanced-layout/api"
	"goal-advanced-layout/internal/handler"
	"goal-advanced-layout/internal/server/middleware"
	"goal-advanced-layout/pkg/config"
	jwt2 "goal-advanced-layout/pkg/jwt"
	"goal-advanced-layout/pkg/log"
	mock_service "goal-advanced-layout/test/mocks/service"
)

var (
	userId = "xxx"
)

var logger *log.Logger
var hdl *handler.Handler
var jwt *jwt2.JWT
var router *gin.Engine

func TestMain(m *testing.M) {
	fmt.Println("begin")

	var envConf = flag.String("conf", "../../../config/dev.yaml", "config path, eg: -conf ./config/local.yml")
	flag.Parse()
	conf := config.NewConfig(*envConf)

	logger = log.NewLog()
	hdl = handler.NewHandler(logger)

	jwt = jwt2.NewJwt(conf.Server)
	gin.SetMode(gin.TestMode)
	router = gin.Default()
	router.Use(
		middleware.CORS(),
		middleware.ResponseLog(logger),
		middleware.RequestLog(logger),
		// middleware.SignMiddleware(log),
	)

	code := m.Run()
	fmt.Println("test end")

	os.Exit(code)
}

func TestUserHandler_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	params := v1.RegisterRequest{
		Username: "xxx",
		Password: "123456",
		Email:    "xxx@gmail.com",
	}

	mockUserService := mock_service.NewMockUserService(ctrl)
	mockUserService.EXPECT().Register(gomock.Any(), &params).Return(nil)

	userHandler := handler.NewUserHandler(hdl, mockUserService)
	router.POST("/register", userHandler.Register)

	paramsJson, _ := json.Marshal(params)

	resp := performRequest(router, "POST", "/register", bytes.NewBuffer(paramsJson))

	assert.Equal(t, resp.Code, http.StatusOK)
	// Add assertions for the response body if needed
}

func TestUserHandler_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	params := v1.LoginRequest{
		Username: "xxx",
		Password: "123456",
	}

	mockUserService := mock_service.NewMockUserService(ctrl)
	mockUserService.EXPECT().Login(gomock.Any(), &params).Return("", nil)

	userHandler := handler.NewUserHandler(hdl, mockUserService)
	router.POST("/login", userHandler.Login)
	paramsJson, _ := json.Marshal(params)

	resp := performRequest(router, "POST", "/login", bytes.NewBuffer(paramsJson))

	assert.Equal(t, resp.Code, http.StatusOK)
	// Add assertions for the response body if needed
}

func TestUserHandler_GetProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mock_service.NewMockUserService(ctrl)
	mockUserService.EXPECT().GetProfile(gomock.Any(), userId).Return(&v1.GetProfileResponseData{
		UserId:   userId,
		Username: "xxxxx",
		Nickname: "xxxxx",
	}, nil)

	userHandler := handler.NewUserHandler(hdl, mockUserService)
	router.Use(middleware.NoStrictAuth(jwt, logger))
	router.GET("/user", userHandler.GetProfile)
	req, _ := http.NewRequest("GET", "/user", nil)
	req.Header.Set("Authorization", "Bearer "+genToken(t))

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, http.StatusOK)
	// Add assertions for the response body if needed
}

func TestUserHandler_UpdateProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	params := v1.UpdateProfileRequest{
		Nickname: "alan",
		Email:    "alan@gmail.com",
		Avatar:   "xxx",
	}

	mockUserService := mock_service.NewMockUserService(ctrl)
	mockUserService.EXPECT().UpdateProfile(gomock.Any(), userId, &params).Return(nil)

	userHandler := handler.NewUserHandler(hdl, mockUserService)
	router.Use(middleware.StrictAuth(jwt, logger))
	router.PUT("/user", userHandler.UpdateProfile)
	paramsJson, _ := json.Marshal(params)

	req, _ := http.NewRequest("PUT", "/user", bytes.NewBuffer(paramsJson))
	req.Header.Set("Authorization", "Bearer "+genToken(t))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusOK)
	// Add assertions for the response body if needed
}

func performRequest(r http.Handler, method, path string, body *bytes.Buffer) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	return resp
}

func genToken(t *testing.T) string {
	token, err := jwt.GenToken(userId, time.Now().Add(time.Hour*24*90))
	if err != nil {
		t.Error(err)
		return token
	}
	return token
}
