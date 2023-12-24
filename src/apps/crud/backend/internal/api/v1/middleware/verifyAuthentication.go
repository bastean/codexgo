package middleware

import (
	"net/http"
	"strings"

	"github.com/bastean/codexgo/backend/internal/util/error"
	"github.com/bastean/codexgo/context/pkg/shared/infrastructure/authentication"
	"github.com/gin-gonic/gin"
)

func VerifyAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": error.AuthenticationMissing("Token")})
			return
		}

		value := strings.Split(token, " ")[1]

		claims := authentication.ValidateJWT(value)

		if value, ok := claims["id"]; ok {
			c.Set("id", value)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": error.AuthenticationMissing("Id")})
		}
	}
}
