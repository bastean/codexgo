package auth

import (
	"os"

	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/sauthentication"
)

type Payload sauthentication.Payload

var secretKey = os.Getenv("JWT_SECRET_KEY")

var auth = sauthentication.NewAuthentication(secretKey)

var GenerateJWT = auth.GenerateJWT

var ValidateJWT = auth.ValidateJWT
