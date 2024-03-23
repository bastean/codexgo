package handler

import (
	"net/http"

	"github.com/bastean/codexgo/pkg/cmd/server/event"
	"github.com/bastean/codexgo/pkg/cmd/server/service/user"
	"github.com/bastean/codexgo/pkg/context/user/application/delete"
	"github.com/gin-gonic/gin"
)

type Delete struct {
	Id string `form:"id"`
}

func FormDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userDelete Delete

		c.ShouldBind(&userDelete)

		id, _ := c.Get("id")

		userDelete.Id = id.(string)

		command := delete.Command(userDelete)

		user.UserDeleteHandler.Handle(&command)

		c.Header("HX-Trigger", event.Client.DeleteAuthorization)

		c.Header("HX-Refresh", "true")

		c.Status(http.StatusOK)
	}
}
