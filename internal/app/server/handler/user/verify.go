package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bastean/codexgo/v4/internal/app/server/service/errs"
	"github.com/bastean/codexgo/v4/internal/app/server/service/key"
	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
	"github.com/bastean/codexgo/v4/internal/pkg/service/module/user"
)

func Verify(c *gin.Context) {
	id := c.Param(key.Id)

	if id == "" {
		errs.AbortByErrWithRedirect(c, errs.MissingKey(key.Id, "Verify"), "/")
		return
	}

	command := new(user.VerifyCommand)

	command.Id = id

	err := user.Verify.Handle(command)

	if err != nil {
		errs.AbortByErr(c, errors.BubbleUp(err, "Verify"))
		return
	}

	c.Redirect(http.StatusFound, "/dashboard")
}
