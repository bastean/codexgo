package page

import (
	"net/http"

	"github.com/bastean/codexgo/internal/app/server/component/page/dashboard"
	"github.com/bastean/codexgo/internal/app/server/util/errs"
	"github.com/bastean/codexgo/internal/app/server/util/key"
	"github.com/bastean/codexgo/internal/pkg/service/user"
	"github.com/gin-gonic/gin"
)

func Dashboard() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, exists := c.Get(key.UserId)

		if !exists {
			c.Error(errs.MissingKey(key.UserId, "Dashboard"))
			c.Redirect(http.StatusFound, "/")
			c.Abort()
			return
		}

		query := new(user.ReadQuery)

		query.Id = id.(string)

		user, err := user.Read.Handle(query)

		if err != nil {
			c.Error(err)
			c.Redirect(http.StatusFound, "/")
			c.Abort()
			return
		}

		dashboard.Page(user).Render(c.Request.Context(), c.Writer)
	}
}
