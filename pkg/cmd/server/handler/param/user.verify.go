package param

import (
	"net/http"

	"github.com/bastean/codexgo/pkg/cmd/server/service/user"
	"github.com/bastean/codexgo/pkg/context/user/application/verify"
	"github.com/gin-gonic/gin"
)

func UserVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Param("id") == "" {
			c.Abort()
			return
		}

		command := new(verify.Command)

		command.Id = c.Param("id")

		err := user.VerifyHandler.Handle(command)

		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		c.Redirect(http.StatusFound, "/dashboard")
	}
}
