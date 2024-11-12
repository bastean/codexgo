package api

import (
	"github.com/gin-gonic/gin"

	"github.com/bastean/codexgo/v4/internal/app/server/handler/user"
	"github.com/bastean/codexgo/v4/internal/app/server/middleware"
)

type API struct {
	*gin.RouterGroup
}

func (a *API) Public() {
	public := a.Group("/")

	account := public.Group("/account")

	account.PUT("/", user.Create)
	account.POST("/", user.Login)

	account.GET("/verify/:id", user.Verify)
}

func (a *API) Private() {
	private := a.Group("/", middleware.Authentication)

	account := private.Group("/account")

	account.PATCH("/", user.Update)
	account.DELETE("/", user.Delete)
}

func Use(group *gin.RouterGroup) {
	api := &API{
		RouterGroup: group,
	}

	api.Public()

	api.Private()
}
