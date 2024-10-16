package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bastean/codexgo/v4/internal/app/server/service/errs"
	"github.com/bastean/codexgo/v4/internal/app/server/service/reply"
	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
	"github.com/bastean/codexgo/v4/internal/pkg/service/module/user"
)

func Create(c *gin.Context) {
	command := new(user.CreateCommand)

	err := c.BindJSON(command)

	if err != nil {
		errs.AbortByErr(c, errs.BindingJSON(err, "Create"))
		return
	}

	err = user.Create.Handle(command)

	if err != nil {
		errs.AbortByErr(c, errors.BubbleUp(err, "Create"))
		return
	}

	c.JSON(http.StatusCreated, &reply.JSON{
		Success: true,
		Message: "Account created",
	})
}
