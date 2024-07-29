package router

import (
	"embed"
	"net/http"

	"github.com/bastean/codexgo/v4/internal/app/server/middleware"
	"github.com/bastean/codexgo/v4/internal/pkg/service/env"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func New(files *embed.FS) *gin.Engine {
	gin.SetMode(env.ServerGinMode)

	Router = gin.Default()

	Router.Use(gin.CustomRecovery(middleware.Recover))

	Router.Use(middleware.Error())

	Router.Use(middleware.Headers())

	Router.Use(middleware.RateLimiter())

	Router.Use(middleware.CookieSession())

	fs := http.FS(files)

	Router.StaticFS("/public", fs)

	Router.StaticFileFS("/robots.txt", "static/robots.txt", fs)

	Routes()

	return Router
}
