package handler

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		files, _ := c.Get("embedFiles")

		templates := template.Must(template.ParseFS(files.(embed.FS), "templates/layouts/base.html", "templates/layouts/alert.html", "templates/pages/home.html"))

		c.Status(http.StatusOK)

		templates.ExecuteTemplate(c.Writer, "base", nil)
	}
}
