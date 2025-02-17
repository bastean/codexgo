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
	"github.com/bastean/codexgo/v4/pkg/context/user/application/create"
)

type CreateData struct {
	*create.CommandAttributes
	*captcha.Data
}

func Create(c *gin.Context) {
	data := new(CreateData)

	err := c.BindJSON(data)

	if err != nil {
		errs.AbortByErr(c, errs.BindingJSON(err, "Create"))
		return
	}

	err = captcha.Verify(data.CaptchaID, data.CaptchaAnswer)

	if err != nil {
		errs.AbortByErr(c, errors.BubbleUp(err, "Create"))
		return
	}

	err = command.Bus.Dispatch(messages.New(
		create.CommandKey,
		&create.CommandAttributes{
			Verify:   services.GenerateID(),
			ID:       services.GenerateID(),
			Email:    data.Email,
			Username: data.Username,
			Password: data.Password,
		},
		new(create.CommandMeta),
	))

	if err != nil {
		errs.AbortByErr(c, errors.BubbleUp(err, "Create"))
		return
	}

	captcha.Clear(data.CaptchaID)

	c.JSON(http.StatusCreated, &reply.JSON{
		Success: true,
		Message: "Account created",
	})
}
