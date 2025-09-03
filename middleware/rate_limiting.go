package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func getClientIP(ctx *gin.Context) string {
	ip := ctx.ClientIP()

	return ip
}

func RateLimitingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := getClientIP(ctx)
		log.Println(ip)
	}
}
