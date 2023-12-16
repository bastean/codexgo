package auth

import (
	"net/http"

	"github.com/bastean/codexgo/backend/internal/container"
	"github.com/bastean/codexgo/context/pkg/user/application/update"
	"github.com/gin-gonic/gin"
)

type Patch struct {
	Id              string `json:"id" binding:"required"`
	Email           string `json:"email"`
	Username        string `json:"username"`
	CurrentPassword string `json:"currentPassword"`
	UpdatedPassword string `json:"updatedPassword"`
}

func UserPatch() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user Patch

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		container.UserUpdateHandler.Handle(update.Command(user))

		c.JSON(http.StatusOK, user)
	}
}
