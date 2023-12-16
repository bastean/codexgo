package public

import (
	"net/http"

	"github.com/bastean/codexgo/backend/internal/container"
	"github.com/bastean/codexgo/context/pkg/user/application/login"
	"github.com/gin-gonic/gin"
)

type Post struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func UserPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user Post

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		container.UserLoginHandler.Handle(login.Query(user))

		c.JSON(http.StatusOK, user)
	}
}
