package bcrypt

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

type Bcrypt struct{}

func (*Bcrypt) Hash(plain string) (string, error) {
	salt := 10
	hashed, err := bcrypt.GenerateFromPassword([]byte(plain), salt)

	if err != nil {
		return "", errors.NewInternal(&errors.Bubble{
			Where: "Hash",
			What:  "Failure to generate a hash",
			Who:   err,
		})
	}

	return string(hashed), nil
}

func (*Bcrypt) IsNotEqual(hashed, plain string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain)) != nil
}
