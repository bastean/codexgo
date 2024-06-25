package middleware

import (
	"net/http"
	"strings"

	"github.com/bastean/codexgo/pkg/cmd/server/service/authentication/jwt"
	"github.com/bastean/codexgo/pkg/cmd/server/util/key"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func abort(c *gin.Context) {
	c.Redirect(http.StatusFound, "/")
	c.Abort()
}

func VerifyAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		token := session.Get(key.Authorization)

		if token == nil {
			abort(c)
			return
		}

		signature := strings.Split(token.(string), " ")[1]

		claims, err := jwt.Validate(signature)

		if err != nil {
			c.Error(err)
			abort(c)
			return
		}

		if value, exists := claims[key.UserId]; exists {
			c.Set(key.UserId, value)
			c.Next()
		} else {
			abort(c)
		}
	}
}
