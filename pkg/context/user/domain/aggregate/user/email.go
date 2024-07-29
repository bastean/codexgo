package user

import (
	"strings"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
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
			What:  "Invalid email format",
			Why: errors.Meta{
				"Email": value,
			},
		})
	}

	return valueObj, nil
}
