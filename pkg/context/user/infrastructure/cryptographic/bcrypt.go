package cryptographic

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/serror"
	"golang.org/x/crypto/bcrypt"
)

type Bcrypt struct{}

func (hashing *Bcrypt) Hash(plain string) (string, error) {
	salt := 10
	bytes, err := bcrypt.GenerateFromPassword([]byte(plain), salt)

	if err != nil {
		return "", serror.NewInternal(&serror.Bubble{
			Where: "Hash",
			What:  "failure to generate a hash",
			Who:   err,
		})
	}

	return string(bytes), nil
}

func (hashing *Bcrypt) IsNotEqual(hashed, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))

	return err != nil
}
