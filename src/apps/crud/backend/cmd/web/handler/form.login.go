package handler

import (
	"net/http"

	"github.com/bastean/codexgo/backend/internal/service"
	"github.com/bastean/codexgo/context/pkg/shared/infrastructure/authentication"
	"github.com/bastean/codexgo/context/pkg/user/application/login"
	"github.com/gin-gonic/gin"
)

type Post struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func FormLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user Post

		c.ShouldBind(&user)

		response := service.UserLoginHandler.Handle(login.Query(user))

		token := authentication.GenerateJWT(response.Id)

		c.Header("Authorization", "Bearer "+token)

		c.Redirect(http.StatusFound, "/dashboard")
	}
}
