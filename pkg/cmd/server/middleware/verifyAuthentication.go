package middleware

import (
	"net/http"
	"strings"

	"github.com/bastean/codexgo/pkg/cmd/server/service/auth"
	"github.com/gin-gonic/gin"
)

func abort(c *gin.Context) {
	c.Redirect(http.StatusFound, "/")
	c.Abort()
}

func VerifyAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			abort(c)
			return
		}

		value := strings.Split(token, " ")[1]

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
