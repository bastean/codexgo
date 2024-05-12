package auth

import (
	"os"

	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/sauthentication"
)

var secretKey = os.Getenv("JWT_SECRET_KEY")

var Auth = sauthentication.NewAuthentication(secretKey)
