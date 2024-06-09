package valueobjs

import (
	"strings"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
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

func NewEmail(email string) (models.ValueObject[string], error) {
	email = strings.TrimSpace(email)

	emailVO := &Email{
		Email: email,
	}

	if emailVO.IsValid() != nil {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewEmail",
			What:  "invalid email format",
			Why: errors.Meta{
				"Email": email,
			},
		})
	}

	return emailVO, nil
}
