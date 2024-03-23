package handler

import (
	"net/http"

	"github.com/bastean/codexgo/pkg/cmd/server/component/partial"
	"github.com/bastean/codexgo/pkg/cmd/server/service/user"
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
		var userPatch Patch

		c.ShouldBind(&userPatch)

		id, _ := c.Get("id")

		userPatch.Id = id.(string)

		command := update.Command(userPatch)

		user.UserUpdateHandler.Handle(&command)

		c.Status(http.StatusOK)

		partial.AlertMsg("success", "Successfully Updated").Render(c.Request.Context(), c.Writer)
	}
}
