package events

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

func Random() *Event {
	return &Event{
		ID:         services.Create.UUID(),
		OccurredOn: services.Create.TimeZoneFull(),
		Key:        Key(services.Create.LoremIpsumWord()),
		Attributes: services.Create.LoremIpsumWord(),
		Meta:       services.Create.LoremIpsumWord(),
	}
}

func RandomWithKey(key Key) *Event {
	return &Event{
		ID:         services.Create.UUID(),
		OccurredOn: services.Create.TimeZoneFull(),
		Key:        key,
		Attributes: services.Create.LoremIpsumWord(),
		Meta:       services.Create.LoremIpsumWord(),
	}
}

func RandomWithAttributes(attributes any) *Event {
	if err := services.Create.Struct(attributes); err != nil {
		errors.Panic(err.Error(), "RandomWithAttributes")
	}

	return &Event{
		ID:         services.Create.UUID(),
		OccurredOn: services.Create.TimeZoneFull(),
		Key:        Key(services.Create.LoremIpsumWord()),
		Attributes: attributes,
		Meta:       services.Create.LoremIpsumWord(),
	}
}

func RandomAttributes(attributes any) {
	if err := services.Create.Struct(attributes); err != nil {
		errors.Panic(err.Error(), "RandomAttributes")
	}
}
