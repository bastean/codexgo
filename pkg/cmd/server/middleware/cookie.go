package middleware

import (
	"github.com/bastean/codexgo/pkg/cmd/server/service/env"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var secretKey = env.Cookie.SecretKey

var sessionName = env.Cookie.SessionName

func CookieSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte(secretKey))
	return sessions.Sessions(sessionName, store)
}
