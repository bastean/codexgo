package server

import (
	"embed"
	"html/template"
	"net/http"
	"os"

	"github.com/bastean/codexgo/backend/cmd/web/middleware"
	"github.com/bastean/codexgo/backend/internal/container"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var server = gin.Default()

func Init(files *embed.FS) *gin.Engine {
	container.Logger.Info("starting server")

	server.Use(middleware.RateLimiter())

	templates := template.Must(template.ParseFS(files, "templates/**/*.html"))

	server.SetHTMLTemplate(templates)

	fs := http.FS(files)

	server.StaticFS("/public", fs)

	server.StaticFileFS("/robots.txt", "static/robots.txt", fs)

	server.Use(gin.CustomRecovery(middleware.ErrorHandler))

	store := cookie.NewStore([]byte(os.Getenv("COOKIE_SECRET_KEY")))

	store.Options(sessions.Options{MaxAge: 60 * 60 * 24 * 7})

	server.Use(sessions.Sessions(os.Getenv("COOKIE_SESSION_NAME"), store))

	server.Use(middleware.EmbedFS(files))

	LoadRoutes()

	container.Logger.Info("listening and serving HTTP on :" + os.Getenv("PORT"))

	return server
}
