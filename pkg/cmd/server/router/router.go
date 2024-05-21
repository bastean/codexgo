package router

import (
	"embed"
	"net/http"

	"github.com/bastean/codexgo/pkg/cmd/server/middleware"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func New(files *embed.FS) *gin.Engine {
	router.Use(gin.CustomRecovery(middleware.PanicHandler))

	router.Use(middleware.ErrorHandler())

	router.Use(middleware.SecurityConfig())

	router.Use(middleware.RateLimiter())

	router.Use(middleware.CookieSession())

	fs := http.FS(files)

	router.StaticFS("/public", fs)

	router.StaticFileFS("/robots.txt", "static/robots.txt", fs)

	InitRoutes()

	return router
}
