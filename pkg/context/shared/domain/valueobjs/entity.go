package valueobjs

import (
	"strings"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/go-playground/validator/v10"
)

const EntityMinCharactersLength = "1"
const EntityMaxCharactersLength = "20"

type Entity struct {
	Entity string `validate:"gte=1,lte=20,alpha"`
}

func (value *Entity) Value() string {
	return value.Entity
}

func (value *Entity) IsValid() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(value)
}

func NewEntity(value string) (models.ValueObject[string], error) {
	value = strings.TrimSpace(value)

	valueObj := &Entity{
		Entity: value,
	}

	if valueObj.IsValid() != nil {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewEntity",
			What:  "entity must be between " + EntityMinCharactersLength + " to " + EntityMaxCharactersLength + " characters and be alpha only",
			Why: errors.Meta{
				"Entity": value,
			},
		})
	}

	return valueObj, nil
}
