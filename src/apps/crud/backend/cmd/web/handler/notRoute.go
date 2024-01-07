package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/")
	}
}
