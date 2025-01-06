package user

import (
	"strings"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

type ID struct {
	Value string `validate:"uuid4"`
}

func NewID(value string) (*ID, error) {
	value = strings.TrimSpace(value)

	object := &ID{
		Value: value,
	}

	if services.IsValueObjectInvalid(object) {
		return nil, errors.New[errors.InvalidValue](&errors.Bubble{
			Where: "NewID",
			What:  "Invalid UUID4 format",
			Why: errors.Meta{
				"ID": value,
			},
		})
	}

	return object, nil
}
