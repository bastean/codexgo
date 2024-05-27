package form

import (
	"net/http"

	"github.com/bastean/codexgo/pkg/cmd/server/service/user"
	"github.com/bastean/codexgo/pkg/cmd/server/util/reply"
	"github.com/gin-gonic/gin"
)

func UserCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		command := new(user.CreateCommand)

		c.BindJSON(command)

		err := user.Create.Handle(command)

		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		c.JSON(http.StatusCreated, reply.JSON(true, "account created", reply.Payload{}))
	}
}
