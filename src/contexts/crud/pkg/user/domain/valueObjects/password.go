package valueObjects

import (
	"github.com/bastean/codexgo/context/pkg/shared/domain/errors"
	"github.com/go-playground/validator/v10"
)

type Password struct {
	Value string `validate:"gte=8,lte=64"`
}

const PasswordMinCharactersLength = "8"
const PasswordMaxCharactersLength = "64"

var InvalidPasswordValue = errors.InvalidValue{Message: "Password must be between " + PasswordMinCharactersLength + " to " + PasswordMaxCharactersLength + " characters"}

func ensureIsValidPassword(password *Password) (err error) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	err = validate.Struct(password)

	return
}

func NewPassword(password string) *Password {
	passwordVO := &Password{password}

	err := ensureIsValidPassword(passwordVO)

	if err != nil {
		panic(InvalidPasswordValue)
	}

	return passwordVO
}
