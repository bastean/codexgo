package valueobjs

import (
	"strings"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/go-playground/validator/v10"
)

const EventMinCharactersLength = "1"
const EventMaxCharactersLength = "20"

type Event struct {
	Event string `validate:"gte=1,lte=20,alpha"`
}

func (value *Event) Value() string {
	return value.Event
}

func (value *Event) IsValid() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(value)
}

func NewEvent(value string) (models.ValueObject[string], error) {
	value = strings.TrimSpace(value)

	valueObj := &Entity{
		Entity: value,
	}

	if valueObj.IsValid() != nil {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewEvent",
			What:  "event must be between " + EventMinCharactersLength + " to " + EventMaxCharactersLength + " characters and be alpha only",
			Why: errors.Meta{
				"Event": value,
			},
		})
	}

	return valueObj, nil
}
