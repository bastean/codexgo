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

		if token != "" {
			value := strings.Split(token, " ")[1]

			claims := authentication.ValidateJWT(value)

			if value, ok := claims["id"]; ok {
				c.Set("id", value)
				c.Next()
				return
			}

			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "missing id"})
			return
		}

		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "missing token"})
	}
}
