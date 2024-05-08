package sharedValueObject

import (
	"strings"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errs"
	"github.com/bastean/codexgo/pkg/context/shared/domain/model"
	"github.com/go-playground/validator/v10"
)

type Email struct {
	value string `validate:"email"`
}

func (email *Email) Value() string {
	return email.value
}

func (email *Email) IsValid() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	// TODO!(fix): structs exposed fields
	return validate.Struct(email)
}

func NewEmail(email string) (model.ValueObject[string], error) {
	email = strings.TrimSpace(email)

	emailVO := &Email{
		value: email,
	}

	if emailVO.IsValid() != nil {
		return nil, errs.NewInvalidValueError(&errs.Bubble{
			Where: "NewEmail",
			What:  "invalid format",
			Why: errs.Meta{
				"Email": email,
			},
		})
	}

	return emailVO, nil
}
