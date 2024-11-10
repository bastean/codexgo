package user

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/bastean/codexgo/v4/internal/app/server/service/errs"
	"github.com/bastean/codexgo/v4/internal/app/server/service/format"
	"github.com/bastean/codexgo/v4/internal/app/server/service/key"
	"github.com/bastean/codexgo/v4/internal/app/server/service/reply"
	"github.com/bastean/codexgo/v4/internal/pkg/service/command"
	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
	"github.com/bastean/codexgo/v4/internal/pkg/service/module/user"
)

func Delete(c *gin.Context) {
	id, exists := c.Get(key.UserId)

	if !exists {
		errs.AbortByErr(c, errs.MissingKey(key.UserId, "Delete"))
		return
	}

	attributes := new(user.DeleteCommandAttributes)

	err := c.BindJSON(attributes)

	if err != nil {
		errs.AbortByErr(c, errs.BindingJSON(err, "Delete"))
		return
	}

	attributes.Id = format.ToString(id)

	err = command.Bus.Dispatch(command.New(
		user.DeleteCommandKey,
		attributes,
		new(user.DeleteCommandMeta),
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
