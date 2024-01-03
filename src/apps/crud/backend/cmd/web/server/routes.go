package server

import (
	"github.com/bastean/codexgo/backend/cmd/web/handler"
)

func LoadRoutes() {
	server.GET("/", handler.IndexPage())
	server.PUT("/", handler.FormRegister())
	server.POST("/", handler.FormLogin())

	/*
		auth := server.Group("/dashboard", middleware.VerifyAuthentication())
		auth.GET("/", handler.IndexPage())
		auth.PATCH("/", handler.FormUpdate())
		auth.DELETE("/", handler.FormDelete())
	*/
}
