package api

import (
	"github.com/bastean/codexgo/v4/internal/app/server/handler/health"
	"github.com/bastean/codexgo/v4/internal/app/server/handler/user"
	"github.com/bastean/codexgo/v4/internal/app/server/middleware"
	"github.com/gin-gonic/gin"
)

type API struct {
	*gin.RouterGroup
}

func (api *API) Public() {
	public := api.Group("/")

	account := public.Group("/account")

	account.PUT("/", user.Create())
	account.POST("/", user.Login())

	account.GET("/verify/:id", user.Verify())
}

func (api *API) Private() {
	private := api.Group("/", middleware.Authentication())

	account := private.Group("/account")

	account.PATCH("/", user.Update())
	account.DELETE("/", user.Delete())
}

func Use(router *gin.Engine) {
	api := &API{
		RouterGroup: router.Group("/v4"),
	}

	router.HEAD("/health", health.Check())

	api.Public()

	api.Private()
}
