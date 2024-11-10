package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bastean/codexgo/v4/internal/app/server/service/errs"
	"github.com/bastean/codexgo/v4/internal/app/server/service/key"
	"github.com/bastean/codexgo/v4/internal/pkg/service/command"
	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
	"github.com/bastean/codexgo/v4/internal/pkg/service/module/user"
)

func Verify(c *gin.Context) {
	id := c.Param(key.Id)

	if id == "" {
		errs.AbortByErrWithRedirect(c, errs.MissingKey(key.Id, "Verify"), "/")
		return
	}

	attributes := new(user.VerifyCommandAttributes)

	attributes.Id = id

	err := command.Bus.Dispatch(command.New(
		user.VerifyCommandKey,
		attributes,
		new(user.VerifyCommandMeta),
	))

	if err != nil {
		errs.AbortByErr(c, errors.BubbleUp(err, "Verify"))
		return
	}

	c.Redirect(http.StatusFound, "/dashboard")
}
