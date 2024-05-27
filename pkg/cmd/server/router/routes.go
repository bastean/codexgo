package router

import (
	"github.com/bastean/codexgo/pkg/cmd/server/handler/page"
	"github.com/bastean/codexgo/pkg/cmd/server/handler/user"
	"github.com/bastean/codexgo/pkg/cmd/server/middleware"
)

func InitRoutes() {
	router.NoRoute(page.Default())

	public := router.Group("/")

	public.GET("/", page.Home())
	public.PUT("/", user.Create())
	public.POST("/", user.Login())

	public.GET("/verify/:id", user.Verify())

	auth := public.Group("/dashboard", middleware.VerifyAuthentication())

	auth.GET("/", page.Dashboard())
	auth.PATCH("/", user.Update())
	auth.DELETE("/", user.Delete())
}
