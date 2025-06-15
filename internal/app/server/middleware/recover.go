package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/bastean/codexgo/v4/internal/app/server/service/reply"
	"github.com/bastean/codexgo/v4/internal/pkg/adapter/log"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

func Recover() gin.RecoveryFunc {
	return func(c *gin.Context, err any) {
		failure, ok := err.(error)

		if !ok {
			failure = errors.New[errors.Internal](&errors.Bubble{
				What: "Unknown Error",
				Why: errors.Meta{
					"Error": err,
				},
			})
		}

		log.Error(failure.Error())

		c.Abort()

		reply.FailureServer(c)
	}
}
