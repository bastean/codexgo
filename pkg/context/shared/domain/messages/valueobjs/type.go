package valueobjs

import (
	"strings"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/services"
)

var TypeOneOf = []string{"event", "command"}

type Type struct {
	Value string `validate:"oneof=event command"`
}

func NewType(value string) (*Type, error) {
	value = strings.TrimSpace(value)

	valueObj := &Type{
		Value: value,
	}

	if services.IsValueObjectInvalid(valueObj) {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewType",
			What:  "type must be only one of these values: " + strings.Join(TypeOneOf, ", "),
			Why: errors.Meta{
				"Type": value,
			},
		})
	}

	return valueObj, nil
}
