package user

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/bastean/codexgo/v4/internal/app/server/service/errs"
	"github.com/bastean/codexgo/v4/internal/app/server/service/format"
	"github.com/bastean/codexgo/v4/internal/app/server/service/key"
	"github.com/bastean/codexgo/v4/internal/app/server/service/reply"
	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
	"github.com/bastean/codexgo/v4/internal/pkg/service/module/user"
)

func Delete(c *gin.Context) {
	id, exists := c.Get(key.UserId)

	if !exists {
		errs.AbortByErr(c, errs.MissingKey(key.UserId, "Delete"))
		return
	}

	command := new(user.DeleteCommand)

	err := c.BindJSON(command)

	if err != nil {
		errs.AbortByErr(c, errs.BindingJSON(err, "Delete"))
		return
	}

	command.Id = format.ToString(id)

	err = user.Delete.Handle(command)

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
