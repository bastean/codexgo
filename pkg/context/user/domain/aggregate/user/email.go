package user

import (
	"strings"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/services"
)

type Email struct {
	Value string `validate:"email"`
}

func NewEmail(value string) (*Email, error) {
	value = strings.TrimSpace(value)

	valueObj := &Email{
		Value: value,
	}

	if services.IsValueObjectInvalid(valueObj) {
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
