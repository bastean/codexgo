package page

import (
	"github.com/bastean/codexgo/pkg/cmd/server/component/layout"
	"github.com/bastean/codexgo/pkg/cmd/server/component/page"
	"github.com/gin-gonic/gin"
)

func Dashboard() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("HX-Push-Url", "/dashboard")

		// TODO!: page.Dashboard(false)
		layout.Base(page.Dashboard(false)).Render(c.Request.Context(), c.Writer)
	}
}
