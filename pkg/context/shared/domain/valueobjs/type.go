package valueobjs

import (
	"strings"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/go-playground/validator/v10"
)

var TypeOneOf = []string{"event", "command"}

type Type struct {
	Type string `validate:"oneof=event command"`
}

func (value *Type) Value() string {
	return value.Type
}

func (value *Type) IsValid() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(value)
}

func NewType(value string) (models.ValueObject[string], error) {
	value = strings.TrimSpace(value)

	valueObj := &Type{
		Type: value,
	}

	if valueObj.IsValid() != nil {
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
