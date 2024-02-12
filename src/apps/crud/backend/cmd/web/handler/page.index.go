package handler

import (
	"github.com/bastean/codexgo/backend/cmd/web/components/layouts"
	"github.com/bastean/codexgo/backend/cmd/web/components/pages"
	"github.com/gin-gonic/gin"
)

func IndexPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("HX-Push-Url", "/")
		layouts.Base(pages.Home()).Render(c.Request.Context(), c.Writer)
	}
}
