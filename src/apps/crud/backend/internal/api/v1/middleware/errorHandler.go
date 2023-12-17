package middleware

import (
	"net/http"

	"github.com/bastean/codexgo/context/pkg/shared/domain/errors"
	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context, err any) {
	switch error := err.(type) {
	case errors.InvalidValue:
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, error.Message)
	case errors.NotExist:
		c.AbortWithStatusJSON(http.StatusNotFound, error.Message)
	case errors.AlreadyExist:
		c.AbortWithStatusJSON(http.StatusConflict, error.Message)
	default:
		c.AbortWithStatusJSON(500, error)
	}
}
