package middleware

import (
	"net/http"

	"github.com/bastean/codexgo/backend/cmd/web/components/partials"
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

	c.Status(code)

	partials.AlertMsg("error", err.(error).Error()).Render(c.Request.Context(), c.Writer)

	c.Abort()
}
