package handler

import (
	"net/http"

	"github.com/bastean/codexgo/backend/internal/container"
	"github.com/bastean/codexgo/context/pkg/user/application/update"
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

		container.UserUpdateHandler.Handle(update.Command(user))

		c.HTML(http.StatusOK, "alert.msg.html", gin.H{"Type": "success", "Message": "Successfully Updated"})
	}
}
