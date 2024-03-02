package handler

import (
	"net/http"

	"github.com/bastean/codexgo/pkg/cmd/server/component/partial"
	"github.com/bastean/codexgo/pkg/cmd/server/service"
	"github.com/bastean/codexgo/pkg/context/user/application/update"
	"github.com/gin-gonic/gin"
)

type Patch struct {
	Id              string `form:"id"`
	Email           string `form:"email"`
	Username        string `form:"username"`
	Password        string `form:"password"`
	UpdatedPassword string `form:"updatedPassword"`
}

func FormUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user Patch

		c.ShouldBind(&user)

		id, _ := c.Get("id")

		user.Id = id.(string)

		service.UserUpdateHandler.Handle(update.Command(user))

		c.Status(http.StatusOK)

		partial.AlertMsg("success", "Successfully Updated").Render(c.Request.Context(), c.Writer)
	}
}
