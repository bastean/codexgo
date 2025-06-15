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
	"github.com/bastean/codexgo/v4/pkg/context/user/application/create"
)

type CreateData struct {
	*create.CommandAttributes
	*captcha.Data
}

func Create(c *gin.Context) {
	data := new(CreateData)

	err := c.ShouldBindJSON(data)

	if err != nil {
		errs.AbortByErr(c, errs.BindingJSON(err))
		return
	}

	if data.Data == nil {
		errs.AbortByErr(c, errs.Missing("Captcha"))
		return
	}

	err = captcha.Verify(data.CaptchaID, data.CaptchaAnswer)

	if err != nil {
		errs.AbortByErr(c, errors.BubbleUp(err))
		return
	}

	if data.CommandAttributes == nil {
		errs.AbortByErr(c, errs.Missing("Attributes"))
		return
	}

	err = command.Bus.Dispatch(messages.New(
		create.CommandKey,
		&create.CommandAttributes{
			VerifyToken: id.New(),
			ID:          id.New(),
			Email:       data.Email,
			Username:    data.Username,
			Password:    data.Password,
		},
		new(create.CommandMeta),
	))

	if err != nil {
		errs.AbortByErr(c, errors.BubbleUp(err))
		return
	}

	captcha.Clear(data.CaptchaID)

	reply.Success(c, "Account created")
}
