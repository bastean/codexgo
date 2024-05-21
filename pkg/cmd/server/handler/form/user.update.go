package form

import (
	"net/http"

	"github.com/bastean/codexgo/pkg/cmd/server/service/user"
	"github.com/bastean/codexgo/pkg/cmd/server/util/reply"
	"github.com/bastean/codexgo/pkg/context/user/application/update"
	"github.com/gin-gonic/gin"
)

func UserUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, exists := c.Get("userId")

		if !exists {
			c.Abort()
			return
		}

		command := new(update.Command)

		c.BindJSON(command)

		command.Id = id.(string)

		err := user.UpdateHandler.Handle(command)

		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, reply.JSON(true, "Account updated", reply.Payload{}))
	}
}
