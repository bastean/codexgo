package auth

import (
	"os"

	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/authentications"
)

type Payload authentications.Payload

var secretKey = os.Getenv("JWT_SECRET_KEY")

var auth = authentications.NewAuthentication(secretKey)

var GenerateJWT = auth.GenerateJWT

var ValidateJWT = auth.ValidateJWT
