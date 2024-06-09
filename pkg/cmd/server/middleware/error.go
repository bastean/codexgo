package middleware

import (
	"net/http"

	"github.com/bastean/codexgo/pkg/cmd/server/service/errors"
	"github.com/bastean/codexgo/pkg/cmd/server/service/logger"
	"github.com/bastean/codexgo/pkg/cmd/server/util/reply"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		var invalidValue *errors.InvalidValue
		var alreadyExist *errors.AlreadyExist
		var notExist *errors.NotExist
		var failure *errors.Failure
		var internal *errors.Internal

		for _, err := range c.Errors {
			switch {
			case errors.As(err, &invalidValue):
				c.JSON(http.StatusUnprocessableEntity, reply.JSON(false, invalidValue.What, invalidValue.Why))
			case errors.As(err, &alreadyExist):
				c.JSON(http.StatusConflict, reply.JSON(false, alreadyExist.What, alreadyExist.Why))
			case errors.As(err, &notExist):
				c.JSON(http.StatusNotFound, reply.JSON(false, notExist.What, notExist.Why))
			case errors.As(err, &failure):
				c.JSON(http.StatusBadRequest, reply.JSON(false, failure.What, failure.Why))
			case errors.As(err, &internal):
				c.JSON(http.StatusInternalServerError, reply.JSON(false, "internal server error", reply.Payload{}))
				fallthrough
			default:
				logger.Error(err.Error())
			}
		}
	}
}
