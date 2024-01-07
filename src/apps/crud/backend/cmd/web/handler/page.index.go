package handler

import (
	"github.com/bastean/codexgo/backend/internal/util"
	"github.com/gin-gonic/gin"
)

func IndexPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("HX-Push-Url", "/")

		util.RunEmbedTemplatePage(c,
			"templates/layouts/base.html",
			"templates/layouts/alert.html",
			"templates/pages/home.html",
		)
	}
}
