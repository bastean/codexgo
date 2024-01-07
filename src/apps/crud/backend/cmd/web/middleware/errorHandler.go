package middleware

import (
	"net/http"

	"github.com/bastean/codexgo/context/pkg/shared/domain/errors"
	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context, err any) {
	var code int

	switch err.(type) {
	case errors.InvalidValue:
		code = http.StatusUnprocessableEntity
	case errors.NotExist:
		code = http.StatusNotFound
	case errors.AlreadyExist:
		code = http.StatusConflict
	default:
		code = http.StatusInternalServerError
	}

	c.HTML(code, "alert.msg.html", gin.H{"Type": "error", "Message": err.(error).Error()})

	c.Abort()
}
