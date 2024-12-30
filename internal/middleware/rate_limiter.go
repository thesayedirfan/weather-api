package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thesayedirfan/weather/internal/cache"
)

func RateLimiter(cache *cache.Request, limit int) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()

		if cache.Get(clientIP) >= limit {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded",
			})
			c.Abort()
			return
		}

		cache.Increment(clientIP)

		go func() {
			time.Sleep(1 * time.Minute)
			cache.Decrement(clientIP)
		}()

		c.Next()

	}
}
