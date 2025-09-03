package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type Client struct {
	limiter  *rate.Limiter
	lassSeen time.Time
}

var (
	mu      sync.Mutex
	clients = make(map[string]Client)
)

func getClientIP(ctx *gin.Context) string {
	ip := ctx.ClientIP()
	if ip == "" {
		ip = ctx.Request.RemoteAddr
	}

	return ip
}

func getRateLimiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()
	client, exists := clients[ip]
	// IP does not exist → create new
	if !exists {
		limiter := rate.NewLimiter(5, 10) // 5 request/s , brust : 10 (max)
		newClient := &Client{limiter, time.Now()}
		clients[ip] = *newClient

		return limiter
	}

	client.lassSeen = time.Now()
	return client.limiter
}

func RateLimitingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := getClientIP(ctx)

		limiter := getRateLimiter(ip)
		if !limiter.Allow() {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error":   "Too many request",
				"message": "Bạn đã gửi quá nhiều request. Hảy thử lại sau",
			})
		}
	}
}
