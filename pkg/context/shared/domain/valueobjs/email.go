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

func (value *Email) Value() string {
	return value.Email
}

func (value *Email) IsValid() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(value)
}

func NewEmail(value string) (models.ValueObject[string], error) {
	value = strings.TrimSpace(value)

	valueObj := &Email{
		Email: value,
	}

	if valueObj.IsValid() != nil {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewEmail",
			What:  "invalid email format",
			Why: errors.Meta{
				"Email": value,
			},
		})
	}

	return valueObj, nil
}
