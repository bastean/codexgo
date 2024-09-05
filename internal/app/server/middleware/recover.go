package middleware

import (
	"net/http"

	"github.com/bastean/codexgo/v4/internal/app/server/util/reply"
	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
	"github.com/bastean/codexgo/v4/internal/pkg/service/logger/log"
	"github.com/gin-gonic/gin"
)

func Recover() gin.RecoveryFunc {
	return func(c *gin.Context, err any) {
		failure, ok := err.(error)

		if !ok {
			failure = errors.NewInternal(&errors.Bubble{
				Where: "Recover",
				What:  "Unknown Error",
				Why: errors.Meta{
					"Error": err,
				},
			})
		}

		log.Error(failure.Error())

		c.AbortWithStatusJSON(http.StatusInternalServerError, &reply.JSON{Message: "Server error. Try again later."})
	}
}
