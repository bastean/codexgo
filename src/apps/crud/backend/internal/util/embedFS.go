package util

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RunEmbedTemplatePage(c *gin.Context, files ...string) {
	embedFS, _ := c.Get("embedFS")
	isHxRequest := c.GetBool("isHxRequest")

	templates := template.Must(template.ParseFS(embedFS.(*embed.FS), files...))

	c.Status(http.StatusOK)

	templates.ExecuteTemplate(c.Writer, "base", gin.H{"isNotHxRequest": !isHxRequest})
}
