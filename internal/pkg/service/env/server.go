package env

import (
	"os"
)

var (
	ServerGinHostname          = os.Getenv("CODEXGO_SERVER_GIN_HOSTNAME")
	ServerGinPort              = os.Getenv("CODEXGO_SERVER_GIN_PORT")
	ServerGinURL               = os.Getenv("CODEXGO_SERVER_GIN_URL")
	ServerGinMode              = os.Getenv("CODEXGO_SERVER_GIN_MODE")
	ServerGinAllowedHosts      = os.Getenv("CODEXGO_SERVER_GIN_ALLOWED_HOSTS")
	ServerGinCookieSecretKey   = os.Getenv("CODEXGO_SERVER_GIN_COOKIE_SECRET_KEY")
	ServerGinCookieSessionName = os.Getenv("CODEXGO_SERVER_GIN_COOKIE_SESSION_NAME")
)

func HasServerGinProxy() (string, bool) {
	proxy := os.Getenv("CODEXGO_DEV_AIR_PROXY_PORT")

	if proxy != "" && proxy != ServerGinPort {
		return proxy, true
	}

	return "", false
}
