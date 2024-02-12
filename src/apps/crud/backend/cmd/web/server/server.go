package server

import (
	"embed"
	"net/http"
	"os"

	"github.com/bastean/codexgo/backend/cmd/web/middleware"
	"github.com/bastean/codexgo/backend/internal/service"
	"github.com/gin-gonic/gin"
)

var server = gin.Default()

func Init(files *embed.FS) *gin.Engine {
	service.Logger.Info("starting server")

	server.Use(middleware.SecurityConfig())

	server.Use(middleware.RateLimiter())

	fs := http.FS(files)

	server.StaticFS("/public", fs)

	server.StaticFileFS("/robots.txt", "static/robots.txt", fs)

	server.Use(gin.CustomRecovery(middleware.ErrorHandler))

	LoadRoutes()

	service.Logger.Info("listening and serving HTTP on :" + os.Getenv("PORT"))

	return server
}
