package valueobj

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/go-playground/validator/v10"
)

const PasswordMinCharactersLength = "8"
const PasswordMaxCharactersLength = "64"

type Password struct {
	Password string `validate:"gte=8,lte=64"`
}

func (value *Password) Value() string {
	return value.Password
}

func (value *Password) IsValid() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(value)
}

func NewPassword(value string) (models.ValueObject[string], error) {
	valueObj := &Password{
		Password: value,
	}

	if valueObj.IsValid() != nil {
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
