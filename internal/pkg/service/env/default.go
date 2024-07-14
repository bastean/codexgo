package env

import (
	"os"
)

var RabbitMQ = &struct {
	URI, Name string
}{
	URI:  os.Getenv("BROKER_RABBITMQ_URI"),
	Name: os.Getenv("BROKER_RABBITMQ_NAME"),
}

var MongoDB = &struct {
	URI, Name string
}{
	URI:  os.Getenv("DATABASE_MONGODB_URI"),
	Name: os.Getenv("DATABASE_MONGODB_NAME"),
}

var SMTP = &struct {
	Host, Port, Username, Password string
}{
	Host:     os.Getenv("CODEXGO_SMTP_HOST"),
	Port:     os.Getenv("CODEXGO_SMTP_PORT"),
	Username: os.Getenv("CODEXGO_SMTP_USERNAME"),
	Password: os.Getenv("CODEXGO_SMTP_PASSWORD"),
}

var JWT = &struct {
	SecretKey string
}{
	SecretKey: os.Getenv("CODEXGO_JWT_SECRET_KEY"),
}

type security struct {
	AllowedHosts string
}

type cookie struct {
	SecretKey, SessionName string
}

type server struct {
	URL, Port, Mode string
	Security        *security
	Cookie          *cookie
}

func (server *server) HasProxy() (string, bool) {
	proxy := os.Getenv("CODEXGO_DEV_AIR_PROXY_PORT")

	if proxy != "" && proxy != server.Port {
		return proxy, true
	}

	return "", false
}

var Server = &server{
	URL:  os.Getenv("CODEXGO_SERVER_URL"),
	Port: os.Getenv("CODEXGO_SERVER_GIN_PORT"),
	Mode: os.Getenv("CODEXGO_SERVER_GIN_MODE"),
	Security: &security{
		AllowedHosts: os.Getenv("CODEXGO_SERVER_GIN_ALLOWED_HOSTS"),
	},
	Cookie: &cookie{
		SecretKey:   os.Getenv("CODEXGO_SERVER_COOKIE_SECRET_KEY"),
		SessionName: os.Getenv("CODEXGO_SERVER_COOKIE_SESSION_NAME"),
	},
}
