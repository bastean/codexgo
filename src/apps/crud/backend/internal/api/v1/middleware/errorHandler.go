package middleware

import (
	"net/http"

	"github.com/bastean/codexgo/context/pkg/shared/domain/errors"
	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context, err any) {
	switch error := err.(type) {
	case errors.InvalidValue:
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": error.Dump})
	case errors.NotExist:
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": error.Dump})
	case errors.AlreadyExist:
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": error.Dump})
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": error})
	}
}
