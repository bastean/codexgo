package user

import (
	"github.com/gin-gonic/gin"

	"github.com/bastean/codexgo/v4/internal/app/server/service/errs"
	"github.com/bastean/codexgo/v4/internal/app/server/service/reply"
	"github.com/bastean/codexgo/v4/internal/pkg/adapter/command"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/verify"
)

func Verify(c *gin.Context) {
	attributes := new(verify.CommandAttributes)

	err := c.ShouldBindJSON(attributes)

	if err != nil {
		errs.AbortByErr(c, errs.BindingJSON(err))
		return
	}

	err = command.Bus.Dispatch(messages.New(
		verify.CommandKey,
		attributes,
		new(verify.CommandMeta),
	))

	if err != nil {
		errs.AbortByErr(c, errors.BubbleUp(err))
		return
	}

	reply.Success(c, "Account verified")
}
