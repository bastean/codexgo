package router

import (
	"github.com/bastean/codexgo/pkg/cmd/server/handler"
	"github.com/bastean/codexgo/pkg/cmd/server/middleware"
)

func InitRoutes() {
	router.NoRoute(handler.NotRoute())

	router.GET("/", handler.IndexPage())
	router.PUT("/", handler.FormRegister())
	router.POST("/", handler.FormLogin())

	auth := router.Group("/dashboard", middleware.VerifyAuthentication())
	auth.GET("/", handler.IndexDashboard())
	auth.PATCH("/", handler.FormUpdate())
	auth.DELETE("/", handler.FormDelete())
}
