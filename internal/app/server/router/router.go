package router

import (
	"embed"
	"net/http"
	"os"

	"github.com/bastean/codexgo/internal/app/server/middleware"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func New(files *embed.FS) *gin.Engine {
	gin.SetMode(os.Getenv("CODEXGO_SERVER_GIN_MODE"))

	router = gin.Default()

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
