package middleware

import (
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func LoggerMiddleware() gin.HandlerFunc {
	logPath := "log/http.log"

	if err := os.MkdirAll(filepath.Dir(logPath), os.ModePerm); err != nil {
		panic(err)
	}

	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	logger := zerolog.New(logFile).With().Timestamp().Logger()

	return func(ctx *gin.Context) {
		statt := time.Now()

		ctx.Next()

		duration := time.Since(statt)

		logEvent := logger.Info()

		logEvent.
			Str("method", ctx.Request.Method).
			Str("path", ctx.Request.URL.Path).
			Str("query", ctx.Request.URL.RawPath).
			Str("client_ip", ctx.ClientIP()).
			Str("user_agent", ctx.Request.UserAgent()). // FireFox, Google, Safari...
			Str("referer", ctx.Request.Referer()).      // Zalo, Fb -> my API
			Str("protocal", ctx.Request.Proto).         // http, https
			Str("host", ctx.Request.Host).
			Str("remote_addr", ctx.Request.RemoteAddr). // Proxy add: 1.1.1.
			Str("request_uri", ctx.Request.RequestURI).
			Int64("content_length", ctx.Request.ContentLength).
			Interface("headers", ctx.Request.Header).
			Int("status_code", ctx.Writer.Status()).
			Int64("duration_ms", duration.Microseconds()).
			Msg("HTTP Request Log")

	}
}
