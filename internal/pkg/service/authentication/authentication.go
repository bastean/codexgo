package authentication

import (
	"github.com/bastean/codexgo/v4/internal/pkg/adapter/log"
	"github.com/bastean/codexgo/v4/internal/pkg/service/env"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/authentications/jwt"
)

var Service = &struct {
	JWT string
}{
	JWT: log.Service("JWT"),
}

var (
	JWT *jwt.JWT
)

func Up() error {
	switch {
	default:
		log.Starting(Service.JWT)

		JWT = jwt.New(env.JWTSecretKey)

		log.Started(Service.JWT)
	}

	return nil
}
