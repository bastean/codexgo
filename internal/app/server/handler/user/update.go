package user

import (
	"net/http"

	"github.com/bastean/codexgo/v4/internal/app/server/util/errs"
	"github.com/bastean/codexgo/v4/internal/app/server/util/format"
	"github.com/bastean/codexgo/v4/internal/app/server/util/key"
	"github.com/bastean/codexgo/v4/internal/app/server/util/reply"
	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
	"github.com/bastean/codexgo/v4/internal/pkg/service/module/user"
	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	id, exists := c.Get(key.UserId)

	if !exists {
		errs.AbortErr(c, errs.MissingKey(key.UserId, "Update"))
		return
	}

	command := new(user.UpdateCommand)

	err := c.BindJSON(command)

	if err != nil {
		errs.AbortErr(c, errs.BindingJSON(err, "Update"))
		return
	}

	command.Id = format.ToString(id)

	err = user.Update.Handle(command)

	if err != nil {
		errs.AbortErr(c, errors.BubbleUp(err, "Update"))
		return
	}

	c.JSON(http.StatusOK, &reply.JSON{
		Success: true,
		Message: "Account updated",
	})
}
