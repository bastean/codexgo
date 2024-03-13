package handler

import (
	"net/http"

	"github.com/bastean/codexgo/pkg/cmd/server/event"
	"github.com/bastean/codexgo/pkg/cmd/server/service"
	"github.com/bastean/codexgo/pkg/context/user/application/delete"
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

		command := delete.Command(user)

		service.UserDeleteHandler.Handle(&command)

		c.Header("HX-Trigger", event.Client.DeleteAuthorization)

		c.Header("HX-Refresh", "true")

		c.Status(http.StatusOK)
	}
}
