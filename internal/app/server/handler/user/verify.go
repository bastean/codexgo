package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bastean/codexgo/v4/internal/app/server/service/errs"
	"github.com/bastean/codexgo/v4/internal/app/server/service/key"
	"github.com/bastean/codexgo/v4/internal/pkg/adapter/command"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/verify"
)

func Verify(c *gin.Context) {
	token := c.Query(key.Token)

	if token == "" {
		errs.AbortByErrWithRedirect(c, errs.MissingKey(key.Token, "Verify"), "/")
		return
	}

	id := c.Query(key.ID)

	if id == "" {
		errs.AbortByErrWithRedirect(c, errs.MissingKey(key.ID, "Verify"), "/")
		return
	}

	attributes := &verify.CommandAttributes{
		Verify: token,
		ID:     id,
	}

	err := command.Bus.Dispatch(messages.New(
		verify.CommandKey,
		attributes,
		new(verify.CommandMeta),
	))

	if err != nil {
		errs.AbortByErr(c, errors.BubbleUp(err, "Verify"))
		return
	}

	c.Redirect(http.StatusFound, "/dashboard")
}
