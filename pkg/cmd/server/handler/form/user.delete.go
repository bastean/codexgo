package form

import (
	"net/http"

	"github.com/bastean/codexgo/pkg/cmd/server/service/user"
	"github.com/bastean/codexgo/pkg/cmd/server/util/reply"
	"github.com/bastean/codexgo/pkg/context/user/application/delete"
	"github.com/gin-gonic/gin"
)

func UserDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, exists := c.Get("userId")

		if !exists {
			c.Abort()
			return
		}

		command := new(delete.Command)

		c.BindJSON(command)

		command.Id = id.(string)

		err := user.DeleteHandler.Handle(command)

		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, reply.JSON(true, "account deleted", reply.Payload{}))
	}
}
