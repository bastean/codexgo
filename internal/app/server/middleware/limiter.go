package middleware

import (
	"net/http"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"

	"github.com/bastean/codexgo/v4/internal/pkg/service/env"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/time"
)

func RateLimiter() gin.HandlerFunc {
	return ratelimit.RateLimiter(
		ratelimit.InMemoryStore(
			&ratelimit.InMemoryOptions{
				Rate:  time.Minute,
				Limit: 120,
				Skip: func(c *gin.Context) bool {
					return env.IsServerGinModeTest()
				},
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
