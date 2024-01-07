package middleware

import (
	"embed"

	"github.com/gin-gonic/gin"
)

func EmbedFS(files *embed.FS) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("embedFS", files)
		c.Set("isHxRequest", c.Request.Header["Hx-Request"])
		c.Next()
	}
}
