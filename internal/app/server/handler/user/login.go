package user

import (
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/bastean/codexgo/v4/internal/app/server/service/errs"
	"github.com/bastean/codexgo/v4/internal/app/server/service/key"
	"github.com/bastean/codexgo/v4/internal/app/server/service/reply"
	"github.com/bastean/codexgo/v4/internal/pkg/service/authentication/jwt"
	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
	"github.com/bastean/codexgo/v4/internal/pkg/service/module/user"
	"github.com/bastean/codexgo/v4/internal/pkg/service/query"
)

func Login(c *gin.Context) {
	attributes := new(user.LoginQueryAttributes)

	err := c.BindJSON(attributes)

	if err != nil {
		errs.AbortByErr(c, errs.BindingJSON(err, "Login"))
		return
	}

	response, err := query.Bus.Ask(query.New(
		user.LoginQueryKey,
		attributes,
		new(user.LoginQueryMeta),
	))

	if err != nil {
		errs.AbortByErr(c, errors.BubbleUp(err, "Login"))
		return
	}

	found, ok := response.Attributes.(*user.LoginResponseAttributes)

	if !ok {
		errs.AbortByErr(c, errs.Assertion("Login"))
		return
	}

	token, err := jwt.Generate(jwt.Payload{
		key.Exp:    time.Now().Add((24 * time.Hour) * 7).Unix(),
		key.UserId: found.Id,
	})

	if err != nil {
		errs.AbortByErr(c, errors.BubbleUp(err, "Login"))
		return
	}

	session := sessions.Default(c)

	session.Set(key.Authorization, "Bearer "+token)

	err = session.Save()

	if err != nil {
		errs.AbortByErr(c, errs.SessionSave(err, "Login"))
		return
	}

	c.JSON(http.StatusOK, &reply.JSON{
		Success: true,
		Message: "Logged in",
	})
}
