package public

import (
	"net/http"

	"github.com/bastean/codexgo/backend/internal/container"
	"github.com/bastean/codexgo/backend/internal/server/util/error"
	"github.com/bastean/codexgo/context/pkg/shared/infrastructure/authentication"
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
			c.JSON(http.StatusBadRequest, gin.H{"error": error.Bind(err.Error())})
			return
		}

		response := container.UserLoginHandler.Handle(login.Query(user))

		token := authentication.GenerateJWT(response.Id)

		c.Header("Authorization", "Bearer "+token)

		c.JSON(http.StatusOK, response)
	}
}
