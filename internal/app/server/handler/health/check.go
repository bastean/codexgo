package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Check() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(http.StatusOK)
	}
}
