package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)


var clients = make(map[string]*rate.Limiter)
var mtx sync.Mutex

func getClientLimiter(ip string) *rate.Limiter {
    mtx.Lock()
    defer mtx.Unlock()

    lim, exists := clients[ip]
    if !exists {
        lim = rate.NewLimiter(rate.Every(time.Second), 30)
        clients[ip] = lim
    }
    return lim
}

func RateLimiterMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        limiter := getClientLimiter(c.ClientIP())

        if !limiter.Allow() {
            c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Too Many Requests"})
            return
        }
        c.Next()
    }
}