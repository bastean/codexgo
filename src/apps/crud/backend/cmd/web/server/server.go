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

	templ := template.Must(template.ParseFS(files, "templates/*.html"))

	server.SetHTMLTemplate(templ)

	server.StaticFS("/public", http.FS(files))

	LoadRoutes()

	return server
}
