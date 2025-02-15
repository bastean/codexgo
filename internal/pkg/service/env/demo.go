package env

import (
	"os"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

func InitDemo() error {
	if err := errors.Join(
		os.Setenv(JWT_SECRET_KEY, "codexgo-demo"),
		os.Setenv(SERVER_GIN_HOSTNAME, "localhost"),
		os.Setenv(SERVER_GIN_PORT, "8080"),
		os.Setenv(SERVER_GIN_URL, "http://localhost:8080"),
		os.Setenv(SERVER_GIN_MODE, "release"),
		os.Setenv(SERVER_GIN_ALLOWED_HOSTS, "localhost:8080"),
		os.Setenv(SERVER_GIN_COOKIE_SECRET_KEY, "codexgo-demo"),
		os.Setenv(SERVER_GIN_COOKIE_SESSION_NAME, "codexgo-demo"),
	); err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			Where: "InitDemo",
			What:  "Failure to set ENV",
			Who:   err,
		})
	}

	return nil
}
