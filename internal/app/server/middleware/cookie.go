package middleware

import (
	"github.com/bastean/codexgo/internal/pkg/service/env"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var secretKey = env.Server.Cookie.SecretKey

var sessionName = env.Server.Cookie.SessionName

func CookieSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte(secretKey))
	return sessions.Sessions(sessionName, store)
}
