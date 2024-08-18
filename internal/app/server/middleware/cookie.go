package middleware

import (
	"github.com/bastean/codexgo/v4/internal/pkg/service/env"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

var CookieSession = sessions.Sessions(
	env.ServerGinCookieSessionName,
	cookie.NewStore([]byte(env.ServerGinCookieSecretKey)),
)
