package middleware

import (
	"strings"

	"github.com/bastean/codexgo/v4/internal/app/server/util/errs"
	"github.com/bastean/codexgo/v4/internal/app/server/util/key"
	"github.com/bastean/codexgo/v4/internal/pkg/service/authentication/jwt"
	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		token := session.Get(key.Authorization)

		if token == nil {
			errs.AbortWithRedirect(c, "/")
			return
		}

		signature := strings.Split(token.(string), " ")[1]

		claims, err := jwt.Validate(signature)

		if err != nil {
			errs.AbortErrWithRedirect(c, errors.BubbleUp(err, "Authentication"), "/")
			return
		}

		if value, exists := claims[key.UserId]; exists {
			c.Set(key.UserId, value)
			c.Next()
		} else {
			errs.AbortErrWithRedirect(c, errs.MissingKey(key.UserId, "Authentication"), "/")
		}
	}
}
