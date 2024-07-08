package jwt

import (
	"github.com/bastean/codexgo/internal/pkg/service/env"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/authentications"
)

type Payload = authentications.Payload

var (
	jwt      = authentications.NewJWT(env.JWT.SecretKey)
	Generate = jwt.Generate
	Validate = jwt.Validate
)