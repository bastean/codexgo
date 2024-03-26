package valueObject

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/go-playground/validator/v10"
)

const PasswordMinCharactersLength = "8"
const PasswordMaxCharactersLength = "64"

var InvalidPasswordValue = errors.InvalidValue{Message: "Password must be between " + PasswordMinCharactersLength + " to " + PasswordMaxCharactersLength + " characters"}

type Password struct {
	Value string `validate:"gte=8,lte=64"`
}

func ensureIsValidPassword(password *Password) error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	return validate.Struct(password)
}

func NewPassword(password string) *Password {
	passwordVO := &Password{password}

	err := ensureIsValidPassword(passwordVO)

	if err != nil {
		panic(InvalidPasswordValue)
	}

	return passwordVO
}
