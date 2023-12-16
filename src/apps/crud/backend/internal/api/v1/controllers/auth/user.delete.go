package auth

import (
	"net/http"

	"github.com/bastean/codexgo/backend/internal/container"
	"github.com/bastean/codexgo/context/pkg/user/application/delete"
	"github.com/gin-gonic/gin"
)

type Delete struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func UserDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user Delete

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		container.UserDeleteHandler.Handle(delete.Command(user))

		c.JSON(http.StatusOK, user)
	}
}
