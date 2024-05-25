package form

import (
	"net/http"
	"time"

	"github.com/bastean/codexgo/pkg/cmd/server/service/auth"
	"github.com/bastean/codexgo/pkg/cmd/server/service/user"
	"github.com/bastean/codexgo/pkg/cmd/server/util/key"
	"github.com/bastean/codexgo/pkg/cmd/server/util/reply"
	"github.com/bastean/codexgo/pkg/context/user/application/login"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func UserLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := new(login.Query)

		c.BindJSON(query)

		user, err := user.LoginHandler.Handle(query)

		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		token, err := auth.GenerateJWT(auth.Payload{
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
