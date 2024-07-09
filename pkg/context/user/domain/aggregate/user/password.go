package user

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/services"
)

const PasswordMinCharactersLength = "8"
const PasswordMaxCharactersLength = "64"

type Password struct {
	Value string `validate:"gte=8,lte=64"`
}

func NewPassword(value string) (*Password, error) {
	valueObj := &Password{
		Value: value,
	}

	if services.IsValueObjectInvalid(valueObj) {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewPassword",
			What:  "password must be between " + PasswordMinCharactersLength + " to " + PasswordMaxCharactersLength + " characters",
			Why: errors.Meta{
				"Password": value,
			},
		})
	}

	return valueObj, nil
}
