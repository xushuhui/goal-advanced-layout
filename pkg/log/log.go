package log

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const LOGGER_KEY = "zapLogger"

type Logger struct {
	*zap.Logger
}

func NewLog() *Logger {
	return initZap()
}

func initZap() *Logger {
	// log address "out.log" User-defined
	lp := "out.log"

	var level zapcore.Level
	// debug<info<warn<error<fatal<panic
	level = zap.InfoLevel
	hook := lumberjack.Logger{
		Filename: lp, // Log file path

	}

	var encoder zapcore.Encoder

	encoder = zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // Print to console and file
		level,
	)
	// if conf.GetString("env") != "prod" {
	// 	return &Logger{zap.New(core, zap.Development(), zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
	// }
	return &Logger{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	// enc.AppendString(t.Format("2006-01-02 15:04:05"))
	enc.AppendString(t.Format("2006-01-02 15:04:05.000000000"))
}

// NewContext Adds a field to the specified context
func (l *Logger) NewContext(ctx *gin.Context, fields ...zapcore.Field) {
	ctx.Set(LOGGER_KEY, l.WithContext(ctx).With(fields...))
}

// WithContext Returns a zap instance from the specified context
func (l *Logger) WithContext(ctx *gin.Context) *Logger {
	if ctx == nil {
		return l
	}
	zl, _ := ctx.Get(LOGGER_KEY)
	ctxLogger, ok := zl.(*zap.Logger)
	if ok {
		return &Logger{ctxLogger}
	}
	return l
}
