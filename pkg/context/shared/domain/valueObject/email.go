package sharedValueObject

import (
	"strings"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/go-playground/validator/v10"
)

var InvalidEmailValue = errors.NewInvalidValue("Email Invalid")

type Email struct {
	Value string `validate:"email"`
}

func ensureIsValidEmail(email *Email) error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	return validate.Struct(email)
}

func NewEmail(email string) *Email {
	email = strings.TrimSpace(email)

	emailVO := &Email{email}

	err := ensureIsValidEmail(emailVO)

	if err != nil {
		panic(InvalidEmailValue)
	}

	return emailVO
}
