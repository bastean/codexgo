package router

import (
	"github.com/bastean/codexgo/v4/internal/app/server/handler/health"
	"github.com/bastean/codexgo/v4/internal/app/server/handler/page"
	"github.com/bastean/codexgo/v4/internal/app/server/handler/user"
	"github.com/bastean/codexgo/v4/internal/app/server/middleware"
)

func Routes() {
	Router.NoRoute(page.Default())

	public := Router.Group("/")

	public.HEAD("/", health.Check())

	public.GET("/", page.Home())
	public.PUT("/", user.Create())
	public.POST("/", user.Login())

	public.GET("/verify/:id", user.Verify())

	auth := public.Group("/dashboard", middleware.Authentication())

	auth.GET("/", page.Dashboard())
	auth.PATCH("/", user.Update())
	auth.DELETE("/", user.Delete())
}
