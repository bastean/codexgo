package sharedValueObject

import (
	"strings"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errs"
	"github.com/bastean/codexgo/pkg/context/shared/domain/model"
	"github.com/go-playground/validator/v10"
)

type Email struct {
	Email string `validate:"email"`
}

func (email *Email) Value() string {
	return email.Email
}

func (email *Email) IsValid() error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	return validate.Struct(email)
}

func NewEmail(email string) (model.ValueObject[string], error) {
	email = strings.TrimSpace(email)

	emailVO := &Email{
		Email: email,
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
