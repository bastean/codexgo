package handler

import (
	"net/http"

	"github.com/bastean/codexgo/backend/internal/container"
	"github.com/bastean/codexgo/context/pkg/user/application/delete"
	"github.com/gin-gonic/gin"
)

type Delete struct {
	Id       string `json:"id"`
	Password string `json:"password" binding:"required"`
}

func FormDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user Delete

		if err := c.ShouldBind(&user); err != nil {
			c.HTML(http.StatusBadRequest, "alert-error.html", "Missing values")
		}

		id, _ := c.Get("id")

		user.Id = id.(string)

		container.UserDeleteHandler.Handle(delete.Command(user))

		c.Status(http.StatusOK)
	}
}
