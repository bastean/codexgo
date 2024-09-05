package env

import (
	"os"

	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
)

type (
	key      = string
	required = bool
	what     = any
)

const (
	BROKER_RABBITMQ_URI  = "CODEXGO_BROKER_RABBITMQ_URI"
	BROKER_RABBITMQ_NAME = "CODEXGO_BROKER_RABBITMQ_NAME"
)

const (
	DATABASE_MONGODB_URI  = "CODEXGO_DATABASE_MONGODB_URI"
	DATABASE_MONGODB_NAME = "CODEXGO_DATABASE_MONGODB_NAME"
)

const (
	SMTP_HOST     = "CODEXGO_SMTP_HOST"
	SMTP_PORT     = "CODEXGO_SMTP_PORT"
	SMTP_USERNAME = "CODEXGO_SMTP_USERNAME"
	SMTP_PASSWORD = "CODEXGO_SMTP_PASSWORD"
)

const (
	JWT_SECRET_KEY = "CODEXGO_JWT_SECRET_KEY"
)

const (
	SERVER_GIN_HOSTNAME            = "CODEXGO_SERVER_GIN_HOSTNAME"
	SERVER_GIN_PORT                = "CODEXGO_SERVER_GIN_PORT"
	SERVER_GIN_URL                 = "CODEXGO_SERVER_GIN_URL"
	SERVER_GIN_MODE                = "CODEXGO_SERVER_GIN_MODE"
	SERVER_GIN_ALLOWED_HOSTS       = "CODEXGO_SERVER_GIN_ALLOWED_HOSTS"
	SERVER_GIN_COOKIE_SECRET_KEY   = "CODEXGO_SERVER_GIN_COOKIE_SECRET_KEY"
	SERVER_GIN_COOKIE_SESSION_NAME = "CODEXGO_SERVER_GIN_COOKIE_SESSION_NAME"
)

const (
	DEV_AIR_PROXY_PORT = "CODEXGO_DEV_AIR_PROXY_PORT"
)

var ENV = map[key]required{
	BROKER_RABBITMQ_URI:  true,
	BROKER_RABBITMQ_NAME: true,

	DATABASE_MONGODB_URI:  true,
	DATABASE_MONGODB_NAME: true,

	SMTP_HOST:     false,
	SMTP_PORT:     false,
	SMTP_USERNAME: false,
	SMTP_PASSWORD: false,

	JWT_SECRET_KEY: true,

	SERVER_GIN_HOSTNAME:            true,
	SERVER_GIN_PORT:                true,
	SERVER_GIN_URL:                 true,
	SERVER_GIN_MODE:                true,
	SERVER_GIN_ALLOWED_HOSTS:       true,
	SERVER_GIN_COOKIE_SECRET_KEY:   true,
	SERVER_GIN_COOKIE_SESSION_NAME: true,
}

func Verify() error {
	errs := map[key]what{}

	for key, isRequired := range ENV {
		value, exists := os.LookupEnv(key)

		switch {
		case isRequired && !exists:
			errs[key] = "Missing"
		case isRequired && value == "":
			errs[key] = "Empty"
		}
	}

	if len(errs) > 0 {
		return errors.NewInternal(&errors.Bubble{
			Where: "Verify",
			What:  "Invalid ENV",
			Why:   errs,
		})
	}

	return nil
}

func Init() error {
	if err := Verify(); err != nil {
		return errors.BubbleUp(err, "Init")
	}

	SMTP()

	Broker()

	Database()

	JWT()

	Server()

	return nil
}
