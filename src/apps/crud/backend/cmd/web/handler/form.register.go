package handler

import (
	"net/http"

	"github.com/bastean/codexgo/backend/internal/container"
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

		container.UserRegisterHandler.Handle(register.Command(user))

		c.HTML(http.StatusCreated, "alert.msg.html", gin.H{"Type": "success", "Message": "Successfully Registered"})
	}
}
