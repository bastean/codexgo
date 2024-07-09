package valueobjs

import (
	"strings"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/services"
)

const EventMinCharactersLength = "1"
const EventMaxCharactersLength = "20"

type Event struct {
	Value string `validate:"gte=1,lte=20,alpha"`
}

func NewEvent(value string) (*Event, error) {
	value = strings.TrimSpace(value)

	valueObj := &Event{
		Value: value,
	}

	if services.IsValueObjectInvalid(valueObj) {
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
