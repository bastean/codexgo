package middleware

import (
	"net/http"

	"github.com/bastean/codexgo/v4/internal/app/server/util/key"
	"github.com/bastean/codexgo/v4/internal/app/server/util/reply"
	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
	"github.com/bastean/codexgo/v4/internal/pkg/service/logger/log"
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Status int
	*reply.JSON
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		var (
			errInvalidValue *errors.ErrInvalidValue
			errAlreadyExist *errors.ErrAlreadyExist
			errNotExist     *errors.ErrNotExist
			errFailure      *errors.ErrFailure
			errInternal     *errors.ErrInternal
		)

		if len(c.Errors) == 0 {
			return
		}

		var response *ErrorResponse

		err := c.Errors[len(c.Errors)-1]

		switch {
		case errors.As(err, &errInvalidValue):
			response = &ErrorResponse{http.StatusUnprocessableEntity, &reply.JSON{Message: errInvalidValue.What, Data: errInvalidValue.Why}}
		case errors.As(err, &errAlreadyExist):
			response = &ErrorResponse{http.StatusConflict, &reply.JSON{Message: errAlreadyExist.What, Data: errAlreadyExist.Why}}
		case errors.As(err, &errNotExist):
			response = &ErrorResponse{http.StatusNotFound, &reply.JSON{Message: errNotExist.What, Data: errNotExist.Why}}
		case errors.As(err, &errFailure):
			response = &ErrorResponse{http.StatusBadRequest, &reply.JSON{Message: errFailure.What, Data: errFailure.Why}}
		case errors.As(err, &errInternal):
			response = &ErrorResponse{http.StatusInternalServerError, &reply.JSON{Message: "Server error. Try again later."}}
			fallthrough
		case err != nil:
			log.Error(err.Error())
		}

		if route, shouldRedirect := c.Get(key.Redirect); shouldRedirect {
			if response != nil {
				log.Error(err.Error())
			}

			route, ok := route.(string)

			if !ok {
				log.Error(errors.NewInternal(&errors.Bubble{
					Where: "ErrorHandler",
					What:  "Invalid redirection Route",
					Why: errors.Meta{
						"Route": route,
					},
				}).Error())

				route = "/"
			}

			c.Redirect(http.StatusFound, route)

			return
		}

		if response != nil {
			c.JSON(response.Status, response.JSON)
		}
	}
}
