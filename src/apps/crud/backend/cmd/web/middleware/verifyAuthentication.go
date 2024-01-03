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
		token := session.Get("Authorization").(string)

		if token == "" {
			c.Redirect(http.StatusUnprocessableEntity, "/")
			return
		}

		value := strings.Split(token, " ")[1]

		claims := authentication.ValidateJWT(value)

		if value, ok := claims["id"]; ok {
			c.Set("id", value)
			c.Next()
		} else {
			c.Redirect(http.StatusUnprocessableEntity, "/")
		}
	}
}
