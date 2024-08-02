package middleware

import (
	"net/http"

	"github.com/bastean/codexgo/v4/internal/app/server/util/reply"
	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
	"github.com/bastean/codexgo/v4/internal/pkg/service/logger/log"
	"github.com/gin-gonic/gin"
)

func Error() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		var (
			errInvalidValue *errors.ErrInvalidValue
			errAlreadyExist *errors.ErrAlreadyExist
			errNotExist     *errors.ErrNotExist
			errFailure      *errors.ErrFailure
			errInternal     *errors.ErrInternal
		)

		for _, err := range c.Errors {
			switch {
			case errors.As(err, &errInvalidValue):
				c.JSON(http.StatusUnprocessableEntity, &reply.JSON{Message: errInvalidValue.What, Data: errInvalidValue.Why})
			case errors.As(err, &errAlreadyExist):
				c.JSON(http.StatusConflict, &reply.JSON{Message: errAlreadyExist.What, Data: errAlreadyExist.Why})
			case errors.As(err, &errNotExist):
				c.JSON(http.StatusNotFound, &reply.JSON{Message: errNotExist.What, Data: errNotExist.Why})
			case errors.As(err, &errFailure):
				c.JSON(http.StatusBadRequest, &reply.JSON{Message: errFailure.What, Data: errFailure.Why})
			case errors.As(err, &errInternal):
				c.JSON(http.StatusInternalServerError, &reply.JSON{Message: "Server error. Try again later."})
				fallthrough
			default:
				log.Error(err.Error())
			}
		}
	}
}
