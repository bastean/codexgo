package auth

import (
	"github.com/bastean/codexgo/pkg/cmd/server/service/env"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/authentications"
)

type Payload authentications.Payload

var auth = authentications.NewAuthentication(env.JWT.SecretKey)

var GenerateJWT = auth.GenerateJWT

var ValidateJWT = auth.ValidateJWT
