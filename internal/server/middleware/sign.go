package middleware

import (
	"net/http"
	"sort"
	"strings"

	v1 "goal-advanced-layout/api"
	"goal-advanced-layout/pkg/helper/md5"
	"goal-advanced-layout/pkg/log"

	"github.com/gin-gonic/gin"
)

func SignMiddleware(logger *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requiredHeaders := []string{"Timestamp", "Nonce", "Sign", "App-Version"}

		for _, header := range requiredHeaders {
			value, ok := ctx.Request.Header[header]
			if !ok || len(value) == 0 {
				v1.Fail(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
				ctx.Abort()
				return
			}
		}

		data := map[string]string{
			"AppKey":     "your app key",
			"Timestamp":  ctx.Request.Header.Get("Timestamp"),
			"Nonce":      ctx.Request.Header.Get("Nonce"),
			"AppVersion": ctx.Request.Header.Get("App-Version"),
		}

		var keys []string
		for k := range data {
			keys = append(keys, k)
		}
		sort.Slice(keys, func(i, j int) bool { return strings.ToLower(keys[i]) < strings.ToLower(keys[j]) })

		var str string
		for _, k := range keys {
			str += k + data[k]
		}
		if ctx.Request.Header.Get("Sign") != strings.ToUpper(md5.Md5(str)) {
			v1.Fail(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
