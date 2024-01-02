package server

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/bastean/codexgo/backend/internal/container"
	"github.com/gin-gonic/gin"
)

var server = gin.Default()

func Init(files *embed.FS) *gin.Engine {
	container.Logger.Info("starting server")

	templates := template.Must(template.ParseFS(files, "templates/*.html"))

	server.SetHTMLTemplate(templates)

	fs := http.FS(files)

	server.StaticFS("/public", fs)

	server.StaticFileFS("/robots.txt", "static/robots.txt", fs)

	LoadRoutes()

	return server
}
