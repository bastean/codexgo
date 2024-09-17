package router

import (
	"embed"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bastean/codexgo/v4/internal/app/server/handler/health"
	"github.com/bastean/codexgo/v4/internal/app/server/handler/redirect"
	"github.com/bastean/codexgo/v4/internal/app/server/middleware"
	"github.com/bastean/codexgo/v4/internal/app/server/router/api"
	"github.com/bastean/codexgo/v4/internal/app/server/router/view"
	"github.com/bastean/codexgo/v4/internal/pkg/service/env"
)

var (
	Router *gin.Engine
)

func New(files *embed.FS) *gin.Engine {
	gin.SetMode(env.ServerGinMode)

	Router = gin.Default()

	Router.Use(gin.CustomRecovery(middleware.Recover()))

	Router.Use(middleware.ErrorHandler())

	Router.Use(middleware.SecureHeaders())

	Router.Use(middleware.RateLimiter())

	Router.Use(middleware.CookieSession())

	fs := http.FS(files)

	Router.StaticFS("/public", fs)

	Router.StaticFileFS("/robots.txt", "static/robots.txt", fs)

	api.Use(Router.Group("/v4"))

	view.Use(Router)

	Router.HEAD("/health", health.Check)

	Router.NoRoute(redirect.Default)

	return Router
}
