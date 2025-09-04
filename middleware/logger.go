package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
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
		start := time.Now()
		contentType := ctx.GetHeader("Content-Type")
		requestBody := make(map[string]any)

		// Content-Type: multipart/form-data
		if strings.HasPrefix(contentType, "multipart/form-data") {
			log.Println("multipart/form-data")
		} else {

			bodyBytes, err := io.ReadAll(ctx.Request.Body)
			if err != nil {
				logger.Error().Err(err).Msg("Failed to read request body")
			}

			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

			// Content-Type: application/json
			if strings.HasPrefix(contentType, "application/json") {
				_ = json.Unmarshal(bodyBytes, &requestBody)
			} else {
				// Content-Type: application/x-www-form-urlencoded
				values, _ := url.ParseQuery(string(bodyBytes))
				for key, vals := range values {
					if len(vals) == 1 {
						requestBody[key] = vals[0]
					} else {
						requestBody[key] = vals
					}
				}
			}
		}

		ctx.Next()

		duration := time.Since(start)

		statusCode := ctx.Writer.Status()

		logEvent := logger.Info()
		if statusCode >= 500 {
			logEvent = logger.Error()
		} else if statusCode >= 400 {
			logEvent = logger.Warn()
		}

		logEvent.
			Str("method", ctx.Request.Method).
			Str("path", ctx.Request.URL.Path).
			Str("query", ctx.Request.URL.RawQuery).
			Str("client_ip", ctx.ClientIP()).
			Str("user_agent", ctx.Request.UserAgent()). // FireFox, Google, Safari, Postman...
			Str("referer", ctx.Request.Referer()).      // Zalo, Fb -> my API
			Str("protocol", ctx.Request.Proto).         // http, https
			Str("host", ctx.Request.Host).
			Str("remote_addr", ctx.Request.RemoteAddr). // Proxy address: 1.1.1.
			Str("request_uri", ctx.Request.RequestURI).
			Int64("content_length", ctx.Request.ContentLength).
			Interface("headers", ctx.Request.Header).
			Interface("request_body", requestBody).
			Int("status_code", statusCode).
			Int64("duration_ms", duration.Microseconds()).
			Msg("HTTP Request Log")

	}
}
