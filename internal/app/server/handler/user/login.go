package user

import (
	"net/http"
	"time"

	"github.com/bastean/codexgo/v4/internal/app/server/util/errs"
	"github.com/bastean/codexgo/v4/internal/app/server/util/key"
	"github.com/bastean/codexgo/v4/internal/app/server/util/reply"
	"github.com/bastean/codexgo/v4/internal/pkg/service/authentication/jwt"
	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
	"github.com/bastean/codexgo/v4/internal/pkg/service/user"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := new(user.LoginQuery)

		err := c.BindJSON(query)

		if err != nil {
			errs.AbortErr(c, errs.BindingJSON(err, "Login"))
			return
		}

		found, err := user.Login.Handle(query)

		if err != nil {
			errs.AbortErr(c, errors.BubbleUp(err, "Login"))
			return
		}

		token, err := jwt.Generate(jwt.Payload{
			key.Exp:    time.Now().Add((24 * time.Hour) * 7).Unix(),
			key.UserId: found.Id,
		})

		if err != nil {
			errs.AbortErr(c, errors.BubbleUp(err, "Login"))
			return
		}

		session := sessions.Default(c)

		session.Set(key.Authorization, "Bearer "+token)

		err = session.Save()

		if err != nil {
			errs.AbortErr(c, errs.SessionSave(err, "Login"))
			return
		}

		c.JSON(http.StatusOK, &reply.JSON{
			Success: true,
			Message: "Logged in",
		})
	}
}
