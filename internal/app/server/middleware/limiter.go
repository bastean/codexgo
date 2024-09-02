package middleware

import (
	"net/http"
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
)

func RateLimiter() gin.HandlerFunc {
	return ratelimit.RateLimiter(
		ratelimit.InMemoryStore(
			&ratelimit.InMemoryOptions{
				Rate:  time.Second,
				Limit: 10,
			},
		),
		&ratelimit.Options{
			ErrorHandler: func(c *gin.Context, info ratelimit.Info) {
				c.Status(http.StatusTooManyRequests)
			},
			KeyFunc: func(c *gin.Context) string {
				return c.ClientIP()
			},
		},
	)
}
