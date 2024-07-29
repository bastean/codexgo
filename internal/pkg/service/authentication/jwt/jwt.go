package jwt

import (
	"github.com/bastean/codexgo/v4/internal/pkg/service/env"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/authentications/jwt"
)

type Payload = jwt.Payload

var (
	JWT      = jwt.New(env.JWTSecretKey)
	Generate = JWT.Generate
	Validate = JWT.Validate
)
