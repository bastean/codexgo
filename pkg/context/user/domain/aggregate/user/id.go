package user

import (
	"strings"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

type Id struct {
	Value string `validate:"uuid4"`
}

func NewId(value string) (*Id, error) {
	value = strings.TrimSpace(value)

	valueObj := &Id{
		Value: value,
	}

	if services.IsValueObjectInvalid(valueObj) {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewId",
			What:  "Invalid UUID4 format",
			Why: errors.Meta{
				"Id": value,
			},
		})
	}

	return valueObj, nil
}
