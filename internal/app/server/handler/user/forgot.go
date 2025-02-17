package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bastean/codexgo/v4/internal/app/server/service/captcha"
	"github.com/bastean/codexgo/v4/internal/app/server/service/errs"
	"github.com/bastean/codexgo/v4/internal/app/server/service/reply"
	"github.com/bastean/codexgo/v4/internal/pkg/adapter/command"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/forgot"
)

type ForgotData struct {
	*forgot.CommandAttributes
	*captcha.Data
}

func Forgot(c *gin.Context) {
	data := new(ForgotData)

	err := c.BindJSON(data)

	if err != nil {
		errs.AbortByErr(c, errs.BindingJSON(err, "Forgot"))
		return
	}

	err = captcha.Verify(data.CaptchaID, data.CaptchaAnswer)

	if err != nil {
		errs.AbortByErr(c, errors.BubbleUp(err, "Forgot"))
		return
	}

	err = command.Bus.Dispatch(messages.New(
		forgot.CommandKey,
		&forgot.CommandAttributes{
			Reset: services.GenerateID(),
			Email: data.Email,
		},
		new(forgot.CommandMeta),
	))

	if err != nil {
		errs.AbortByErr(c, errors.BubbleUp(err, "Forgot"))
		return
	}

	captcha.Clear(data.CaptchaID)

	c.JSON(http.StatusOK, &reply.JSON{
		Success: true,
		Message: "Link sent. Please check your inbox",
	})
}
