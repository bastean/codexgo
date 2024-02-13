package router

import (
	"embed"
	"net/http"

	"github.com/bastean/codexgo/backend/cmd/web/middleware"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func New(files *embed.FS) *gin.Engine {
	router.Use(middleware.SecurityConfig())

	router.Use(middleware.RateLimiter())

	fs := http.FS(files)

	router.StaticFS("/public", fs)

	router.StaticFileFS("/robots.txt", "static/robots.txt", fs)

	router.Use(gin.CustomRecovery(middleware.ErrorHandler))

	InitRoutes()

	return router
}
