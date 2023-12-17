package valueObjects

import (
	"strings"

	"github.com/bastean/codexgo/context/pkg/shared/domain/errors"
	"github.com/go-playground/validator/v10"
)

type Email struct {
	Value string `validate:"email"`
}

var InvalidEmailValue = errors.InvalidValue{Message: "Email value is invalid"}

func ensureIsValidEmail(email *Email) (err error) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	err = validate.Struct(email)

	return
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
