package param

import (
	"net/http"

	"github.com/bastean/codexgo/pkg/cmd/server/service/user"
	"github.com/bastean/codexgo/pkg/cmd/server/util/errs"
	"github.com/bastean/codexgo/pkg/cmd/server/util/key"
	"github.com/gin-gonic/gin"
)

func UserVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param(key.Id)

		if id == "" {
			c.Error(errs.MissingKey(key.Id, "UserVerify"))
			c.Redirect(http.StatusFound, "/")
			c.Abort()
			return
		}

		command := new(user.VerifyCommand)

		command.Id = id

		err := user.Verify.Handle(command)

		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		c.Redirect(http.StatusFound, "/dashboard")
	}
}
