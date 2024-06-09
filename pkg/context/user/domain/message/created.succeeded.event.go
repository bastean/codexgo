package message

import (
	"encoding/json"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/messages"
)

var CreatedSucceededEventTypeRoutingKey = messages.NewRoutingKey(&messages.MessageRoutingKey{
	Module:    "user",
	Version:   "1",
	Type:      messages.Type.Event,
	Aggregate: "user",
	Event:     "created",
	Status:    messages.Status.Succeeded,
})

type CreatedSucceededEventAttributes struct {
	Id, Email, Username string
}

func NewCreatedSucceededEvent(attributes *CreatedSucceededEventAttributes) (*messages.Message, error) {
	attributesJson, err := json.Marshal(attributes)

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "NewCreatedSucceededEvent",
			What:  "failure to create an event message",
			Why: errors.Meta{
				"Routing Key": CreatedSucceededEventTypeRoutingKey,
			},
			Who: err,
		})
	}

	return messages.NewMessage(CreatedSucceededEventTypeRoutingKey, attributesJson, messages.Meta{}), nil
}
