package auth

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

func UserPatch() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user Patch

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id, _ := c.Get("id")

		user.Id = id.(string)

		container.UserUpdateHandler.Handle(update.Command(user))

		c.Status(http.StatusOK)
	}
}
