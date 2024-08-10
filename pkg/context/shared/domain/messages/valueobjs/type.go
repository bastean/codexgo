package valueobjs

import (
	"fmt"
	"strings"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

var (
	TypeOneOf = []string{"Event", "Command"}
)

type Type struct {
	Value string `validate:"oneof=event command"`
}

func NewType(value string) (*Type, error) {
	value = strings.TrimSpace(value)
	value = strings.ToLower(value)

	valueObj := &Type{
		Value: value,
	}

	if services.IsValueObjectInvalid(valueObj) {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewType",
			What:  fmt.Sprintf("Type must be only one of these values: %s", strings.Join(TypeOneOf, ", ")),
			Why: errors.Meta{
				"Type": value,
			},
		})
	}

	return valueObj, nil
}
