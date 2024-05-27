package middleware

import (
	"strings"

	"github.com/bastean/codexgo/pkg/cmd/server/service/env"
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
)

var allowedHosts = env.Security.AllowedHosts

func getAllowedHosts() []string {
	return strings.Split(allowedHosts, ", ")
}

func SecurityConfig() gin.HandlerFunc {
	return secure.New(secure.Config{
		AllowedHosts:         getAllowedHosts(),
		STSSeconds:           315360000,
		STSIncludeSubdomains: true,
		FrameDeny:            true,
		ContentTypeNosniff:   true,
		BrowserXssFilter:     true,
		IENoOpen:             true,
		ReferrerPolicy:       "strict-origin-when-cross-origin",
	})
}
