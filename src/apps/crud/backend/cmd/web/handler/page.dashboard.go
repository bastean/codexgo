package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexDashboard() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "page.index.html", struct{ Auth bool }{true})
	}
}
