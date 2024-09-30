package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bastean/codexgo/v4/internal/app/server/service/errs"
	"github.com/bastean/codexgo/v4/internal/app/server/service/format"
	"github.com/bastean/codexgo/v4/internal/app/server/service/key"
	"github.com/bastean/codexgo/v4/internal/app/server/service/reply"
	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
	"github.com/bastean/codexgo/v4/internal/pkg/service/module/user"
)

func Update(c *gin.Context) {
	id, exists := c.Get(key.UserId)

	if !exists {
		errs.AbortByErr(c, errs.MissingKey(key.UserId, "Update"))
		return
	}

	command := new(user.UpdateCommand)

	err := c.BindJSON(command)

	if err != nil {
		errs.AbortByErr(c, errs.BindingJSON(err, "Update"))
		return
	}

	command.Id = format.ToString(id)

	err = user.Update.Handle(command)

	if err != nil {
		errs.AbortByErr(c, errors.BubbleUp(err, "Update"))
		return
	}

	c.JSON(http.StatusOK, &reply.JSON{
		Success: true,
		Message: "Account updated",
	})
}
