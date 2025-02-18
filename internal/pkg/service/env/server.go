package env

import (
	"os"
)

var (
	ServerGinHostname, ServerGinPort, ServerGinURL, ServerGinMode, ServerGinAllowedHosts, ServerGinCookieSecretKey, ServerGinCookieSessionName string
)

func Server() {
	ServerGinHostname = os.Getenv(SERVER_GIN_HOSTNAME)
	ServerGinPort = os.Getenv(SERVER_GIN_PORT)
	ServerGinURL = os.Getenv(SERVER_GIN_URL)
	ServerGinMode = os.Getenv(SERVER_GIN_MODE)
	ServerGinAllowedHosts = os.Getenv(SERVER_GIN_ALLOWED_HOSTS)
	ServerGinCookieSecretKey = os.Getenv(SERVER_GIN_COOKIE_SECRET_KEY)
	ServerGinCookieSessionName = os.Getenv(SERVER_GIN_COOKIE_SESSION_NAME)
}

func IsServerGinModeTest() bool {
	return ServerGinMode == "test"
}
