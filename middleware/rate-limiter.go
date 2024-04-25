package middleware

import (
	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
)
func RateLimiterMiddleware() gin.HandlerFunc {
	throttle := tollbooth.NewLimiter(1, nil)
	return func(c *gin.Context) {
		if throttle.LimitReached(c.ClientIP()) {
			c.JSON(429, gin.H{"error": "Too Many Requests"})
			c.Abort()
			return
		}
	}
	
}