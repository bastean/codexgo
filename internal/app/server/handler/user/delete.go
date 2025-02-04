package user

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/bastean/codexgo/v4/internal/app/server/service/errs"
	"github.com/bastean/codexgo/v4/internal/app/server/service/format"
	"github.com/bastean/codexgo/v4/internal/app/server/service/key"
	"github.com/bastean/codexgo/v4/internal/app/server/service/reply"
	"github.com/bastean/codexgo/v4/internal/pkg/adapter/command"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/delete"
)

func Delete(c *gin.Context) {
	id, exists := c.Get(key.UserID)

	if !exists {
		errs.AbortByErr(c, errs.MissingKey(key.UserID, "Delete"))
		return
	}

	attributes := new(delete.CommandAttributes)

	err := c.BindJSON(attributes)

	if err != nil {
		errs.AbortByErr(c, errs.BindingJSON(err, "Delete"))
		return
	}

	attributes.ID = format.String(id)

	err = command.Bus.Dispatch(messages.New(
		delete.CommandKey,
		attributes,
		new(delete.CommandMeta),
	))

	if err != nil {
		errs.AbortByErr(c, errors.BubbleUp(err, "Delete"))
		return
	}

	session := sessions.Default(c)

	session.Clear()

	session.Options(sessions.Options{
		Path:   "/",
		MaxAge: -1,
	})

	err = session.Save()

	if err != nil {
		errs.AbortByErr(c, errs.SessionSave(err, "Delete"))
		return
	}

	c.JSON(http.StatusOK, &reply.JSON{
		Success: true,
		Message: "Account deleted",
	})
}
