package env

import (
	"os"
)

var ServerURL = os.Getenv("CODEXGO_SERVER_URL")

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
	Host, Port, Username, Password, ServerURL string
}{
	Host:      os.Getenv("CODEXGO_SMTP_HOST"),
	Port:      os.Getenv("CODEXGO_SMTP_PORT"),
	Username:  os.Getenv("CODEXGO_SMTP_USERNAME"),
	Password:  os.Getenv("CODEXGO_SMTP_PASSWORD"),
	ServerURL: ServerURL,
}

var Security = &struct {
	AllowedHosts string
}{
	AllowedHosts: os.Getenv("CODEXGO_SERVER_GIN_ALLOWED_HOSTS"),
}

var JWT = &struct {
	SecretKey string
}{
	SecretKey: os.Getenv("CODEXGO_JWT_SECRET_KEY"),
}

var Cookie = &struct {
	SecretKey, SessionName string
}{
	SecretKey:   os.Getenv("CODEXGO_SERVER_COOKIE_SECRET_KEY"),
	SessionName: os.Getenv("CODEXGO_SERVER_COOKIE_SESSION_NAME"),
}
