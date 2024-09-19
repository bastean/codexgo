package page

import (
	"github.com/gin-gonic/gin"

	"github.com/bastean/codexgo/v4/internal/app/server/component/page/home"
	"github.com/bastean/codexgo/v4/internal/app/server/util/errs"
)

func Home(c *gin.Context) {
	if err := home.Page().Render(c.Request.Context(), c.Writer); err != nil {
		errs.AbortByErr(c, errs.Render(err, "Home"))
	}
}
