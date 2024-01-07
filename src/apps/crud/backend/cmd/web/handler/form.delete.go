package handler

import (
	"net/http"
	"os"

	"github.com/bastean/codexgo/backend/internal/container"
	"github.com/bastean/codexgo/context/pkg/user/application/delete"
	"github.com/gin-gonic/gin"
)

type Delete struct {
	Id string `form:"id"`
}

func FormDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user Delete

		c.ShouldBind(&user)

		id, _ := c.Get("id")

		user.Id = id.(string)

		container.UserDeleteHandler.Handle(delete.Command(user))

		c.SetCookie(os.Getenv("COOKIE_SESSION_NAME"), "", -1, "/", "localhost", false, true)

		c.Header("HX-Refresh", "true")

		c.Status(http.StatusOK)
	}
}
