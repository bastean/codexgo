package jwt

import (
	"github.com/bastean/codexgo/internal/pkg/service/env"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/authentications/jwt"
)

type Payload = jwt.Payload

var (
	JWT      = jwt.New(env.JWT.SecretKey)
	Generate = JWT.Generate
	Validate = JWT.Validate
)
