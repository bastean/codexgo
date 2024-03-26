package sharedValueObject

import (
	"strings"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/go-playground/validator/v10"
)

var InvalidEmailValue = errors.InvalidValue{Message: "Email Invalid"}

type Email struct {
	Value string `validate:"email"`
}

func ensureIsValidEmail(email *Email) error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	return validate.Struct(email)
}

func NewEmail(email string) *Email {
	email = strings.TrimSpace(email)

	emailVo := &Email{email}

	err := ensureIsValidEmail(emailVo)

	if err != nil {
		panic(InvalidEmailValue)
	}

	return emailVo
}
