package server

import (
	"github.com/bastean/codexgo/backend/cmd/web/handler"
	"github.com/bastean/codexgo/backend/cmd/web/middleware"
)

func LoadRoutes() {
	server.NoRoute(handler.NotRoute())

	server.GET("/", handler.IndexPage())
	server.PUT("/", handler.FormRegister())
	server.POST("/", handler.FormLogin())

	auth := server.Group("/dashboard", middleware.VerifyAuthentication())
	auth.GET("/", handler.IndexDashboard())
	auth.PATCH("/", handler.FormUpdate())
	auth.DELETE("/", handler.FormDelete())
}
