package user

import (
	"github.com/gin-gonic/gin"

	"github.com/bastean/codexgo/v4/internal/app/server/service/errs"
	"github.com/bastean/codexgo/v4/internal/app/server/service/reply"
	"github.com/bastean/codexgo/v4/internal/pkg/adapter/command"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/reset"
)

func Reset(c *gin.Context) {
	attributes := new(reset.CommandAttributes)

	err := c.ShouldBindJSON(attributes)

	if err != nil {
		errs.AbortByErr(c, errs.BindingJSON(err))
		return
	}

	err = command.Bus.Dispatch(messages.New(
		reset.CommandKey,
		attributes,
		new(reset.CommandMeta),
	))

	if err != nil {
		errs.AbortByErr(c, errors.BubbleUp(err))
		return
	}

	reply.Success(c, "Password updated")
}
