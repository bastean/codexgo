package handler

import (
	"encoding/json"

	"github.com/bastean/codexgo/pkg/cmd/server/component/page"
	"github.com/bastean/codexgo/pkg/cmd/server/event"
	"github.com/bastean/codexgo/pkg/cmd/server/service"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/authentication"
	"github.com/bastean/codexgo/pkg/context/user/application/login"
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

		query := login.Query(user)

		response := service.UserLoginHandler.Handle(&query)

		token := authentication.GenerateJWT(response.Id)

		event, err := json.Marshal(map[string]string{event.Client.PutAuthorization: "Bearer " + token})

		if err != nil {
			c.Abort()
			return
		}

		c.Header("HX-Trigger", string(event))

		c.Header("HX-Push-Url", "/dashboard")

		page.Dashboard().Render(c.Request.Context(), c.Writer)
	}
}
