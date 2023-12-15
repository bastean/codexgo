package valueObjects

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type Email struct {
	Value string `validate:"email"`
}

func ensureIsValidEmail(email *Email) (err error) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	err = validate.Struct(email)

	return
}

func NewEmail(email string) (*Email, error) {
	email = strings.TrimSpace(email)
	emailVo := &Email{email}

	err := ensureIsValidEmail(emailVo)

	if err != nil {
		return new(Email), err
	}

	return emailVo, nil
}
