package user

import (
	"net/http"

	"github.com/bastean/codexgo/internal/app/server/util/errs"
	"github.com/bastean/codexgo/internal/app/server/util/reply"
	"github.com/bastean/codexgo/internal/pkg/service/user"
	"github.com/gin-gonic/gin"
)

func Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		command := new(user.CreateCommand)

		err := c.BindJSON(command)

		if err != nil {
			c.Error(errs.BindingJSON(err, "Create"))
			c.Abort()
			return
		}

		err = user.Create.Handle(command)

		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		c.JSON(http.StatusCreated, reply.JSON(true, "Account created", reply.Payload{}))
	}
}
