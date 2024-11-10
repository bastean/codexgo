package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bastean/codexgo/v4/internal/app/server/service/errs"
	"github.com/bastean/codexgo/v4/internal/app/server/service/reply"
	"github.com/bastean/codexgo/v4/internal/pkg/service/command"
	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
	"github.com/bastean/codexgo/v4/internal/pkg/service/module/user"
)

func Create(c *gin.Context) {
	attributes := new(user.CreateCommandAttributes)

	err := c.BindJSON(attributes)

	if err != nil {
		errs.AbortByErr(c, errs.BindingJSON(err, "Create"))
		return
	}

	err = command.Bus.Dispatch(command.New(
		user.CreateCommandKey,
		attributes,
		new(user.CreateCommandMeta),
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
