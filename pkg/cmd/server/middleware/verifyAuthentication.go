package middleware

import (
	"net/http"
	"strings"

	"github.com/bastean/codexgo/pkg/cmd/server/service/auth"
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

		token := session.Get("Authorization")

		if token == nil {
			abort(c)
			return
		}

		value := strings.Split(token.(string), " ")[1]

		claims, err := auth.Auth.ValidateJWT(value)

		if err != nil {
			c.Error(err)
			abort(c)
			return
		}

		if value, exists := claims["userId"]; exists {
			c.Set("userId", value)
			c.Next()
		} else {
			abort(c)
		}
	}
}
