package handler

import (
	"net/http"

	"github.com/bastean/codexgo/backend/cmd/web/components/partials"
	"github.com/bastean/codexgo/backend/internal/service"
	"github.com/bastean/codexgo/context/pkg/user/application/register"
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
		var user Put

		c.ShouldBind(&user)

		user.Id = uuid.NewString()

		service.UserRegisterHandler.Handle(register.Command(user))

		c.Status(http.StatusCreated)

		partials.AlertMsg("success", "Successfully Registered").Render(c.Request.Context(), c.Writer)
	}
}
