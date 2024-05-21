package page

import (
	"github.com/bastean/codexgo/pkg/cmd/server/component/page/dashboard"
	"github.com/bastean/codexgo/pkg/cmd/server/service/user"
	"github.com/bastean/codexgo/pkg/context/user/application/read"
	"github.com/gin-gonic/gin"
)

func Dashboard() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, exists := c.Get("userId")

		if !exists {
			c.Abort()
			return
		}

		query := &read.Query{
			Id: id.(string),
		}

		user, err := user.ReadHandler.Handle(query)

		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		dashboard.Page(user).Render(c.Request.Context(), c.Writer)
	}
}
