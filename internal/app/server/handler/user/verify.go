package user

import (
	"net/http"

	"github.com/bastean/codexgo/v4/internal/app/server/util/errs"
	"github.com/bastean/codexgo/v4/internal/app/server/util/key"
	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
	"github.com/bastean/codexgo/v4/internal/pkg/service/user"
	"github.com/gin-gonic/gin"
)

func Verify() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param(key.Id)

		if id == "" {
			errs.AbortErrWithRedirect(c, errs.MissingKey(key.Id, "Verify"), "/")
			return
		}

		command := new(user.VerifyCommand)

		command.Id = id

		err := user.Verify.Handle(command)

		if err != nil {
			errs.AbortErr(c, errors.BubbleUp(err, "Verify"))
			return
		}

		c.Redirect(http.StatusFound, "/dashboard")
	}
}
