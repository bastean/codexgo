package handler

import (
	"encoding/json"

	"github.com/bastean/codexgo/backend/cmd/web/components/pages"
	"github.com/bastean/codexgo/backend/internal/event"
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

		event, err := json.Marshal(map[string]string{event.Client.PutAuthorization: "Bearer " + token})

		if err != nil {
			c.Abort()
			return
		}

		c.Header("HX-Trigger", string(event))

		c.Header("HX-Push-Url", "/dashboard")

		pages.Dashboard().Render(c.Request.Context(), c.Writer)
	}
}
