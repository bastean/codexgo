package user

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/bastean/codexgo/v4/internal/app/server/service/errs"
	"github.com/bastean/codexgo/v4/internal/app/server/service/key"
	"github.com/bastean/codexgo/v4/internal/app/server/service/reply"
	"github.com/bastean/codexgo/v4/internal/pkg/adapter/query"
	"github.com/bastean/codexgo/v4/internal/pkg/service/authentication"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/time"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/authentications/jwt"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/login"
)

func Login(c *gin.Context) {
	attributes := new(login.QueryAttributes)

	err := c.ShouldBindJSON(attributes)

	if err != nil {
		errs.AbortByErr(c, errs.BindingJSON(err))
		return
	}

	response, err := query.Bus.Ask(messages.New(
		login.QueryKey,
		attributes,
		new(login.QueryMeta),
	))

	if err != nil {
		errs.AbortByErr(c, errors.BubbleUp(err))
		return
	}

	found, ok := response.Attributes.(*login.ResponseAttributes)

	if !ok {
		errs.AbortByErr(c, errs.Assertion())
		return
	}

	token, err := authentication.JWT.Generate(jwt.Payload{
		key.Exp:    time.Now().Add((24 * time.Hour) * 7).Unix(),
		key.UserID: found.ID,
	})

	if err != nil {
		errs.AbortByErr(c, errors.BubbleUp(err))
		return
	}

	session := sessions.Default(c)

	session.Set(key.Authorization, "Bearer "+token)

	err = session.Save()

	if err != nil {
		errs.AbortByErr(c, errs.SessionSave(err))
		return
	}

	reply.Success(c, "Logged in")
}
