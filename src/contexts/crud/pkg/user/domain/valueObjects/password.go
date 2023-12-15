package valueObjects

import "github.com/go-playground/validator/v10"

type Password struct {
	Value string `validate:"gte=8,lte=64"`
}

func ensureIsValidPassword(password *Password) (err error) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	err = validate.Struct(password)

	return
}

func NewPassword(password string) (*Password, error) {
	passwordVO := &Password{password}

	err := ensureIsValidPassword(passwordVO)

	if err != nil {
		return new(Password), err
	}

	return passwordVO, nil
}
