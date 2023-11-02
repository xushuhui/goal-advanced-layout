package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Succeed(ctx *gin.Context, data any) {
	if data == nil {
		data = map[string]any{}
	}
	resp := Response{Code: ErrorCodeMap[ErrSuccess], Message: ErrSuccess.Error(), Data: data}
	ctx.JSON(http.StatusOK, resp)
}

func Fail(ctx *gin.Context, httpCode int, err error, data any) {
	if data == nil {
		data = map[string]string{}
	}
	resp := Response{Code: ErrorCodeMap[err], Message: err.Error(), Data: data}
	ctx.JSON(httpCode, resp)
}

type Error struct {
	Code    int
	Message string
}

var ErrorCodeMap = map[error]int{}

func newError(code int, msg string) error {
	err := errors.New(msg)
	ErrorCodeMap[err] = code
	return err
}

func (e Error) Error() string {
	return e.Message
}
