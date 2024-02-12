package handler

import (
	"net/http"

	"github.com/bastean/codexgo/backend/internal/service"
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

		service.UserDeleteHandler.Handle(delete.Command(user))

		c.Header("HX-Refresh", "true")

		c.Status(http.StatusOK)
	}
}
