package middleware

import (
	"strings"

	"github.com/bastean/codexgo/internal/app/server/util/errs"
	"github.com/bastean/codexgo/internal/app/server/util/key"
	"github.com/bastean/codexgo/internal/pkg/service/authentication/jwt"
	"github.com/bastean/codexgo/internal/pkg/service/errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		token := session.Get(key.Authorization)

		if token == nil {
			errs.AbortWithRedirect(c, errs.MissingKey(key.Authorization, "Authentication"), "/")
			return
		}

		signature := strings.Split(token.(string), " ")[1]

		claims, err := jwt.Validate(signature)

		if err != nil {
			errs.AbortWithRedirect(c, errors.BubbleUp(err, "Authentication"), "/")
			return
		}

		if value, exists := claims[key.UserId]; exists {
			c.Set(key.UserId, value)
			c.Next()
		} else {
			errs.AbortWithRedirect(c, errs.MissingKey(key.UserId, "Authentication"), "/")
		}
	}
}
