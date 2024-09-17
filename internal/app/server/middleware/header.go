package middleware

import (
	"strings"

	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"

	"github.com/bastean/codexgo/v4/internal/pkg/service/env"
)

func SecureHeaders() gin.HandlerFunc {
	return secure.New(secure.Config{
		AllowedHosts:         strings.Split(env.ServerGinAllowedHosts, ","),
		STSSeconds:           315360000,
		STSIncludeSubdomains: true,
		FrameDeny:            true,
		ContentTypeNosniff:   true,
		BrowserXssFilter:     true,
		IENoOpen:             true,
		ReferrerPolicy:       "strict-origin-when-cross-origin",
	})
}
