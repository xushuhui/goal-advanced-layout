package log

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
)

const ctxLoggerKey = "slogLogger"

type Logger struct {
	*slog.Logger
}

func NewLog() *Logger {
	return &Logger{slog.New(slog.NewJSONHandler(os.Stdout, nil))}
}

func (l *Logger) WithValue(ctx context.Context, fields ...slog.Attr) context.Context {
	if c, ok := ctx.(*gin.Context); ok {
		ctx = c.Request.Context()
		c.Request = c.Request.WithContext(context.WithValue(ctx, ctxLoggerKey, l.WithContext(ctx).With(fields)))
		return c
	}
	return context.WithValue(ctx, ctxLoggerKey, l.WithContext(ctx).With(fields))
}

func (l *Logger) WithContext(ctx context.Context) *Logger {
	if c, ok := ctx.(*gin.Context); ok {
		ctx = c.Request.Context()
	}
	zl := ctx.Value(ctxLoggerKey)
	ctxLogger, ok := zl.(*slog.Logger)
	if ok {
		return &Logger{ctxLogger}
	}
	return l
}

func (l *Logger) Errorf(format string, a ...any) {
	l.Error(fmt.Errorf(format, a...).Error())
}
