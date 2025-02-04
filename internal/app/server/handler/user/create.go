package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bastean/codexgo/v4/internal/app/server/service/errs"
	"github.com/bastean/codexgo/v4/internal/app/server/service/reply"
	"github.com/bastean/codexgo/v4/internal/pkg/adapter/command"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/create"
)

func Create(c *gin.Context) {
	attributes := new(create.CommandAttributes)

	err := c.BindJSON(attributes)

	if err != nil {
		errs.AbortByErr(c, errs.BindingJSON(err, "Create"))
		return
	}

	attributes.Verify = services.GenerateID()

	attributes.ID = services.GenerateID()

	err = command.Bus.Dispatch(messages.New(
		create.CommandKey,
		attributes,
		new(create.CommandMeta),
	))

	if err != nil {
		errs.AbortByErr(c, errors.BubbleUp(err, "Create"))
		return
	}

	c.JSON(http.StatusCreated, &reply.JSON{
		Success: true,
		Message: "Account created",
	})
}
