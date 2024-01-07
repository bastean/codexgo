package middleware

import (
	"net/http"
	"strings"

	"github.com/bastean/codexgo/context/pkg/shared/infrastructure/authentication"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func VerifyAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		token := session.Get("Authorization")

		if token == nil || token == "" {
			c.Redirect(http.StatusFound, "/")
			c.Abort()
			return
		}

		value := strings.Split(token.(string), " ")[1]

		claims := authentication.ValidateJWT(value)

		if value, ok := claims["id"]; ok {
			c.Set("id", value)
			c.Next()
		} else {
			c.Redirect(http.StatusFound, "/")
			c.Abort()
		}
	}
}
