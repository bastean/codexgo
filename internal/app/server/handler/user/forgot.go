package user

import (
	"github.com/gin-gonic/gin"

	"github.com/bastean/codexgo/v4/internal/app/server/service/captcha"
	"github.com/bastean/codexgo/v4/internal/app/server/service/errs"
	"github.com/bastean/codexgo/v4/internal/app/server/service/reply"
	"github.com/bastean/codexgo/v4/internal/pkg/adapter/command"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/id"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/forgot"
)

type ForgotData struct {
	*forgot.CommandAttributes
	*captcha.Data
}

func Forgot(c *gin.Context) {
	data := new(ForgotData)

	err := c.ShouldBindJSON(data)

	if err != nil {
		errs.AbortByErr(c, errs.BindingJSON(err))
		return
	}

	err = captcha.Verify(data.CaptchaID, data.CaptchaAnswer)

	if err != nil {
		errs.AbortByErr(c, errors.BubbleUp(err))
		return
	}

	err = command.Bus.Dispatch(messages.New(
		forgot.CommandKey,
		&forgot.CommandAttributes{
			ResetToken: id.New(),
			Email:      data.Email,
		},
		new(forgot.CommandMeta),
	))

	if err != nil {
		errs.AbortByErr(c, errors.BubbleUp(err))
		return
	}

	captcha.Clear(data.CaptchaID)

	reply.Success(c, "Link sent. Please check your inbox")
}
