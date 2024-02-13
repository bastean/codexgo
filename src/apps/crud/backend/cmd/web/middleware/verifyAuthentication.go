package middleware

import (
	"net/http"
	"strings"

	"github.com/bastean/codexgo/context/pkg/shared/infrastructure/authentication"
	"github.com/gin-gonic/gin"
)

func VerifyAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			abort(c)
			return
		}

		value := strings.Split(token, " ")[1]

		claims := authentication.ValidateJWT(value)

		if value, ok := claims["id"]; ok {
			c.Set("id", value)
			c.Next()
		} else {
			abort(c)
		}
	}
}

func abort(c *gin.Context) {
	c.Redirect(http.StatusFound, "/")
	c.Abort()
}
