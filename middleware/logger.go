package middleware

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	logPath := "log/http.log"

	if err := os.MkdirAll(filepath.Dir(logPath), os.ModePerm); err != nil {
		panic(err)
	}

	return func(ctx *gin.Context) {

	}
}
