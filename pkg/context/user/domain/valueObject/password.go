package valueObject

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errs"
	"github.com/bastean/codexgo/pkg/context/shared/domain/model"
	"github.com/go-playground/validator/v10"
)

const PasswordMinCharactersLength = "8"
const PasswordMaxCharactersLength = "64"

type Password struct {
	value string `validate:"gte=8,lte=64"`
}

func (password *Password) Value() string {
	return password.value
}

func (password *Password) IsValid() error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	return validate.Struct(password)
}

func NewPassword(password string) (model.ValueObject[string], error) {
	passwordVO := &Password{
		value: password,
	}

	if passwordVO.IsValid() != nil {
		return nil, errs.NewInvalidValueError(&errs.Bubble{
			Where: "NewPassword",
			What:  "must be between " + PasswordMinCharactersLength + " to " + PasswordMaxCharactersLength + " characters",
			Why: errs.Meta{
				"Password": password,
			},
		})
	}

	return passwordVO, nil
}
