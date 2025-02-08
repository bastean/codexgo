package page

import (
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/bastean/codexgo/v4/internal/app/server/component/page/home"
	"github.com/bastean/codexgo/v4/internal/app/server/service/captcha"
	"github.com/bastean/codexgo/v4/internal/app/server/service/errs"
)

func Home(c *gin.Context) {
	register, errRegister := captcha.Generate()
	forgot, errForgot := captcha.Generate()

	if err := errors.Join(errRegister, errForgot); err != nil {
		errs.AbortByErr(c, errs.Render(err, "Home"))
	}

	if err := home.Page(register, forgot).Render(c.Request.Context(), c.Writer); err != nil {
		errs.AbortByErr(c, errs.Render(err, "Home"))
	}
}
