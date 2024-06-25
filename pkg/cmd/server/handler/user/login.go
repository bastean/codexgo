package user

import (
	"net/http"
	"time"

	"github.com/bastean/codexgo/pkg/cmd/server/service/authentication/jwt"
	"github.com/bastean/codexgo/pkg/cmd/server/service/user"
	"github.com/bastean/codexgo/pkg/cmd/server/util/errs"
	"github.com/bastean/codexgo/pkg/cmd/server/util/key"
	"github.com/bastean/codexgo/pkg/cmd/server/util/reply"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := new(user.LoginQuery)

		err := c.BindJSON(query)

		if err != nil {
			c.Error(errs.BindingJSON(err, "Login"))
			c.Abort()
			return
		}

		user, err := user.Login.Handle(query)

		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		token, err := jwt.Generate(jwt.Payload{
			key.Exp:    time.Now().Add((24 * time.Hour) * 7).Unix(),
			key.UserId: user.Id,
		})

		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		session := sessions.Default(c)

		session.Set(key.Authorization, "Bearer "+token)

		session.Save()

		c.JSON(http.StatusOK, reply.JSON(true, "logged in", reply.Payload{}))
	}
}
