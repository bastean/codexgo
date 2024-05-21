package middleware

import (
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var secretKey = os.Getenv("COOKIE_SECRET_KEY")

var sessionName = os.Getenv("COOKIE_SESSION_NAME")

func CookieSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte(secretKey))
	return sessions.Sessions(sessionName, store)
}
