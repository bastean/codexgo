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

func (password *Password) Value() string {
	return password.Password
}

func (password *Password) IsValid() error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	return validate.Struct(password)
}

func NewPassword(password string) (models.ValueObject[string], error) {
	passwordVO := &Password{
		Password: password,
	}

	if passwordVO.IsValid() != nil {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewPassword",
			What:  "password must be between " + PasswordMinCharactersLength + " to " + PasswordMaxCharactersLength + " characters",
			Why: errors.Meta{
				"Password": password,
			},
		})
	}

	return passwordVO, nil
}
