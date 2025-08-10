package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bastean/codexgo/v4/internal/app/server/service/key"
	"github.com/bastean/codexgo/v4/internal/app/server/service/reply"
	"github.com/bastean/codexgo/v4/internal/pkg/adapter/log"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		if route, shouldRedirect := c.Get(key.Redirect); shouldRedirect {
			for _, err := range c.Errors.Errors() {
				log.Error(err)
			}

			route, ok := route.(string)

			if !ok {
				log.Error(errors.New[errors.Internal](&errors.Bubble{
					What: "Invalid redirection Route",
					Why: errors.Meta{
						"Route": route,
					},
				}).Error())

				route = "/"
			}

			c.Redirect(http.StatusFound, route)

			return
		}

		bubbles := new(errors.Bubbles)

		for _, err := range c.Errors {
			errors.Unwrap(err, bubbles)
		}

		switch {
		case len(bubbles.Internal) > 0:
			ids := make([]string, 0, len(bubbles.Internal))

			for _, internal := range bubbles.Internal {
				log.Error(internal.Error())
				ids = append(ids, internal.ID)
			}

			reply.FailureServer(c, ids...)
		case bubbles.Amount > 0:
			reply.FailureClient(c, reply.Errors(bubbles))
		}
	}
}
