package handler

import (
	"net/http"

	"github.com/bastean/codexgo/backend/internal/container"
	"github.com/bastean/codexgo/context/pkg/shared/infrastructure/authentication"
	"github.com/bastean/codexgo/context/pkg/user/application/login"
	"github.com/gin-gonic/gin"
)

type Post struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func FormLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user Post

		if err := c.ShouldBind(&user); err != nil {
			c.HTML(http.StatusBadRequest, "alert-error.html", "Missing values")
		}

		response := container.UserLoginHandler.Handle(login.Query(user))

		token := authentication.GenerateJWT(response.Id)

		c.Header("Authorization", "Bearer "+token)

		c.JSON(http.StatusOK, response)
	}
}
