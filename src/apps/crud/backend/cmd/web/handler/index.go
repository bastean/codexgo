package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", time.Now().Format(time.UnixDate))
	}
}
