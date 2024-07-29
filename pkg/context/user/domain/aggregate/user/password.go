package user

import (
	"fmt"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
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
			What:  fmt.Sprintf("Password must be between %s to %s characters", PasswordMinCharactersLength, PasswordMaxCharactersLength),
			Why: errors.Meta{
				"Password": value,
			},
		})
	}

	return valueObj, nil
}
