package components

import (
	"fmt"
	"strings"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

var (
	TypeOneOf = []string{"Event", "Command", "Query", "Response"}
)

type Type struct {
	Value string `validate:"oneof=event command query response"`
}

func NewType(value string) (*Type, error) {
	value = strings.TrimSpace(value)

	value = strings.ToLower(value)

	object := &Type{
		Value: value,
	}

	if services.IsValueObjectInvalid(object) {
		return nil, errors.New[errors.InvalidValue](&errors.Bubble{
			Where: "NewType",
			What:  fmt.Sprintf("Type must be only one of these values: %s", strings.Join(TypeOneOf, ", ")),
			Why: errors.Meta{
				"Type": value,
			},
		})
	}

	return object, nil
}
