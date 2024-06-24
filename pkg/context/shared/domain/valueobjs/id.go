package valueobjs

import (
	"strings"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/go-playground/validator/v10"
)

type Id struct {
	Id string `validate:"uuid4"`
}

func (value *Id) Value() string {
	return value.Id
}

func (value *Id) IsValid() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(value)
}

func NewId(value string) (models.ValueObject[string], error) {
	value = strings.TrimSpace(value)

	valueObj := &Id{
		Id: value,
	}

	if valueObj.IsValid() != nil {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewId",
			What:  "invalid uuid4 format",
			Why: errors.Meta{
				"Id": value,
			},
		})
	}

	return valueObj, nil
}
