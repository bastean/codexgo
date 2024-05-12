package page

import (
	"github.com/bastean/codexgo/pkg/cmd/server/component/layout"
	"github.com/bastean/codexgo/pkg/cmd/server/component/page"
	"github.com/gin-gonic/gin"
)

func Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("HX-Push-Url", "/")
		layout.Base(page.Home()).Render(c.Request.Context(), c.Writer)
	}
}
