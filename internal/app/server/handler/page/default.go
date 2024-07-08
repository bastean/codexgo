package page

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Default() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/")
	}
}
