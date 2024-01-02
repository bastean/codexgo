package handler

import (
	"net/http"

	"github.com/bastean/codexgo/backend/internal/container"
	"github.com/bastean/codexgo/context/pkg/user/application/register"
	"github.com/gin-gonic/gin"
)

type Put struct {
	Id       string `form:"id" binding:"required"`
	Email    string `form:"email" binding:"required"`
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func FormRegister() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user Put

		if err := c.ShouldBind(&user); err != nil {
			c.HTML(http.StatusBadRequest, "alert-error.html", "Missing values")
		}

		container.UserRegisterHandler.Handle(register.Command(user))

		c.Status(http.StatusCreated)
	}
}
