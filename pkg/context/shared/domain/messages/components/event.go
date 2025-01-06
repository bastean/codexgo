package components

import (
	"fmt"
	"strings"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

const (
	EventMinCharactersLength = "1"
	EventMaxCharactersLength = "20"
)

type Event struct {
	Value string `validate:"gte=1,lte=20,alpha"`
}

func NewEvent(value string) (*Event, error) {
	value = strings.TrimSpace(value)

	object := &Event{
		Value: value,
	}

	if services.IsValueObjectInvalid(object) {
		return nil, errors.New[errors.InvalidValue](&errors.Bubble{
			Where: "NewEvent",
			What:  fmt.Sprintf("Event must be between %s to %s characters and be alpha only", EventMinCharactersLength, EventMaxCharactersLength),
			Why: errors.Meta{
				"Event": value,
			},
		})
	}

	return object, nil
}
