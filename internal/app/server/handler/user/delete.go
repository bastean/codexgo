package user

import (
	"net/http"

	"github.com/bastean/codexgo/internal/app/server/util/errs"
	"github.com/bastean/codexgo/internal/app/server/util/key"
	"github.com/bastean/codexgo/internal/app/server/util/reply"
	"github.com/bastean/codexgo/internal/pkg/service/user"
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/gin-gonic/gin"
)

func Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, exists := c.Get(key.UserId)

		if !exists {
			errs.Abort(c, errs.MissingKey(key.UserId, "Delete"))
			return
		}

		command := new(user.DeleteCommand)

		err := c.BindJSON(command)

		if err != nil {
			errs.Abort(c, errs.BindingJSON(err, "Delete"))
			return
		}

		command.Id = id.(string)

		err = user.Delete.Handle(command)

		if err != nil {
			errs.Abort(c, errors.BubbleUp(err, "Delete"))
			return
		}

		c.JSON(http.StatusOK, &reply.JSON{
			Success: true,
			Message: "Account deleted",
		})
	}
}
