package handler

import (
	"github.com/bastean/codexgo/pkg/cmd/server/component/layout"
	"github.com/bastean/codexgo/pkg/cmd/server/component/page"
	"github.com/gin-gonic/gin"
)

func IndexDashboard() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("HX-Push-Url", "/dashboard")
		layout.Base(page.Dashboard()).Render(c.Request.Context(), c.Writer)
	}
}
