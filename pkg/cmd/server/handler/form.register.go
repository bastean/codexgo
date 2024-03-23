package handler

import (
	"net/http"

	"github.com/bastean/codexgo/pkg/cmd/server/component/partial"
	"github.com/bastean/codexgo/pkg/cmd/server/service/user"
	"github.com/bastean/codexgo/pkg/context/user/application/register"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Put struct {
	Id       string `form:"id"`
	Email    string `form:"email"`
	Username string `form:"username"`
	Password string `form:"password"`
}

func FormRegister() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userPut Put

		c.ShouldBind(&userPut)

		userPut.Id = uuid.NewString()

		command := register.Command(userPut)

		user.UserRegisterHandler.Handle(&command)

		c.Status(http.StatusCreated)

		partial.AlertMsg("success", "Successfully Registered").Render(c.Request.Context(), c.Writer)
	}
}
