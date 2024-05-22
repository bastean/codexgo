package form

import (
	"net/http"

	"github.com/bastean/codexgo/pkg/cmd/server/service/user"
	"github.com/bastean/codexgo/pkg/cmd/server/util/reply"
	"github.com/bastean/codexgo/pkg/context/user/application/create"
	"github.com/gin-gonic/gin"
)

func UserCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		command := new(create.Command)

		c.BindJSON(command)

		err := user.CreateHandler.Handle(command)

		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		c.JSON(http.StatusCreated, reply.JSON(true, "account created", reply.Payload{}))
	}
}
