package middleware

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/bastean/codexgo/context/pkg/shared/domain/errors"
	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context, err any) {
	files, _ := c.Get("embedFiles")

	templates := template.Must(template.ParseFS(files.(embed.FS), "templates/partials/alert.msg.html"))

	switch err.(type) {
	case errors.InvalidValue:
		c.Status(http.StatusUnprocessableEntity)
	case errors.NotExist:
		c.Status(http.StatusNotFound)
	case errors.AlreadyExist:
		c.Status(http.StatusConflict)
	default:
		c.Status(http.StatusInternalServerError)
	}

	templates.ExecuteTemplate(c.Writer, "alert.msg", err.(error).Error())

	c.Abort()
}
