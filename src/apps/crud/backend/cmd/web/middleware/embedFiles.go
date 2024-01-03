package middleware

import (
	"embed"

	"github.com/gin-gonic/gin"
)

func EmbedFiles(files embed.FS) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("embedFiles", files)
		c.Next()
	}
}
