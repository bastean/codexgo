package handler

import (
	"github.com/bastean/codexgo/backend/internal/util"
	"github.com/gin-gonic/gin"
)

func IndexDashboard() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("HX-Push-Url", "/dashboard")

		util.RunEmbedTemplatePage(c,
			"templates/layouts/base.html",
			"templates/layouts/alert.html",
			"templates/pages/dashboard.html",
		)
	}
}
