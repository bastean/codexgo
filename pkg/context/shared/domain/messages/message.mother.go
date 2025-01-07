package messages

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

func Random() *Message {
	return &Message{
		ID:         services.Create.UUID(),
		OccurredOn: services.Create.TimeZoneFull(),
		Key:        Key(services.Create.LoremIpsumWord()),
		Attributes: services.Create.LoremIpsumWord(),
		Meta:       services.Create.LoremIpsumWord(),
	}
}

func RandomWithKey(key Key) *Message {
	return &Message{
		ID:         services.Create.UUID(),
		OccurredOn: services.Create.TimeZoneFull(),
		Key:        key,
		Attributes: services.Create.LoremIpsumWord(),
		Meta:       services.Create.LoremIpsumWord(),
	}
}

func RandomWithAttributes(attributes any, shouldRandomize bool) *Message {
	if shouldRandomize {
		err := services.Create.Struct(attributes)

		if err != nil {
			errors.Panic(err.Error(), "RandomWithAttributes")
		}
	}

	return &Message{
		ID:         services.Create.UUID(),
		OccurredOn: services.Create.TimeZoneFull(),
		Key:        Key(services.Create.LoremIpsumWord()),
		Attributes: attributes,
		Meta:       services.Create.LoremIpsumWord(),
	}
}

func RandomizeAttributes(attributes any) {
	if err := services.Create.Struct(attributes); err != nil {
		errors.Panic(err.Error(), "RandomAttributes")
	}
}
