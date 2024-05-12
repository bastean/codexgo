package form

import (
	"net/http"

	"github.com/bastean/codexgo/pkg/cmd/server/service/user"
	"github.com/bastean/codexgo/pkg/cmd/server/util/reply"
	"github.com/bastean/codexgo/pkg/context/user/application/register"
	"github.com/gin-gonic/gin"
)

func UserRegister() gin.HandlerFunc {
	return func(c *gin.Context) {
		command := new(register.Command)

		c.BindJSON(command)

		err := user.RegisterHandler.Handle(command)

		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		c.JSON(http.StatusCreated, reply.JSON(true, "Account created", reply.EmptyPayload))
	}
}
