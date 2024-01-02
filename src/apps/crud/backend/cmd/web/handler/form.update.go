package handler

import (
	"net/http"

	"github.com/bastean/codexgo/backend/internal/container"
	"github.com/bastean/codexgo/context/pkg/user/application/update"
	"github.com/gin-gonic/gin"
)

type Patch struct {
	Id              string `json:"id"`
	Email           string `json:"email"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	UpdatedPassword string `json:"updatedPassword"`
}

func FormUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user Patch

		if err := c.ShouldBind(&user); err != nil {
			c.HTML(http.StatusBadRequest, "alert-error.html", "Missing values")
		}

		id, _ := c.Get("id")

		user.Id = id.(string)

		container.UserUpdateHandler.Handle(update.Command(user))

		c.Status(http.StatusOK)
	}
}
