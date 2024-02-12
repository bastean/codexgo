package handler

import (
	"github.com/bastean/codexgo/backend/cmd/web/components/layouts"
	"github.com/bastean/codexgo/backend/cmd/web/components/pages"
	"github.com/gin-gonic/gin"
)

func IndexDashboard() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("HX-Push-Url", "/dashboard")
		layouts.Base(pages.Dashboard()).Render(c.Request.Context(), c.Writer)
	}
}
