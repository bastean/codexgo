package server

import "github.com/bastean/codexgo/backend/cmd/web/handler"

func LoadRoutes() {
	server.GET("/", handler.Index())
	server.POST("/", handler.FormRegister())
}
