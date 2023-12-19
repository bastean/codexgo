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

			authentication.ValidateJWT(value)

			c.Next()
		}

		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "missing token"})
	}
}
