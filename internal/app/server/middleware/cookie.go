package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/bastean/codexgo/v4/internal/pkg/service/env"
)

func CookieSession() gin.HandlerFunc {
	return sessions.Sessions(
		env.ServerGinCookieSessionName,
		cookie.NewStore([]byte(env.ServerGinCookieSecretKey)),
	)
}
