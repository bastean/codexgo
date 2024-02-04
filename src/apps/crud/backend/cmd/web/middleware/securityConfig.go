package middleware

import (
	"os"
	"strings"

	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
)

func getAllowedHosts() []string {
	return strings.Split(os.Getenv("ALLOWED_HOSTS"), ", ")
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
