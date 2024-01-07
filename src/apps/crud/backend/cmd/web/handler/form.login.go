package handler

import (
	"net/http"

	"github.com/bastean/codexgo/backend/internal/container"
	"github.com/bastean/codexgo/context/pkg/shared/infrastructure/authentication"
	"github.com/bastean/codexgo/context/pkg/user/application/login"
	"github.com/gin-contrib/sessions"
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

		response := container.UserLoginHandler.Handle(login.Query(user))

		token := authentication.GenerateJWT(response.Id)

		session := sessions.Default(c)

		session.Set("Authorization", "Bearer "+token)

		session.Save()

		c.Redirect(http.StatusFound, "/dashboard")
	}
}
