package router

import (
	"github.com/bastean/codexgo/pkg/cmd/server/handler/form"
	"github.com/bastean/codexgo/pkg/cmd/server/handler/page"
	"github.com/bastean/codexgo/pkg/cmd/server/handler/param"
	"github.com/bastean/codexgo/pkg/cmd/server/middleware"
)

func InitRoutes() {
	router.NoRoute(page.NoRoute())

	public := router.Group("/")

	public.GET("/", page.Index())
	public.PUT("/", form.UserRegister())
	public.POST("/", form.UserLogin())

	public.GET("/verify/:id", param.UserVerify())

	auth := public.Group("/dashboard", middleware.VerifyAuthentication())

	auth.GET("/", page.Dashboard())
	auth.PATCH("/", form.UserUpdate())
	auth.DELETE("/", form.UserDelete())
}
