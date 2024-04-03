package auth

import (
	"os"

	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/authentication"
)

var secretKey = os.Getenv("JWT_SECRET_KEY")

var Auth = authentication.NewAuthentication(secretKey)
