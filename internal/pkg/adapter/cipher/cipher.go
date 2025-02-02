package cipher

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/ciphers/bcrypt"
)

var (
	Hasher roles.Hasher = new(bcrypt.Bcrypt)
)
