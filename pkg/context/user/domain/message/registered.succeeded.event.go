package message

import (
	"encoding/json"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errs"
	"github.com/bastean/codexgo/pkg/context/shared/domain/message"
)

var RegisteredSucceededEventTypeRoutingKey = message.NewRoutingKey(&message.MessageRoutingKey{
	Module:    "user",
	Version:   "1",
	Type:      message.Type.Event,
	Aggregate: "user",
	Event:     "registered",
	Status:    message.Status.Succeeded,
})

type RegisteredSucceededEventAttributes struct {
	Id       string
	Email    string
	Username string
}

func NewRegisteredSucceededEvent(attributes *RegisteredSucceededEventAttributes) (*message.Message, error) {
	attributesJson, err := json.Marshal(attributes)

	if err != nil {
		return nil, errs.NewInternalError(&errs.Bubble{
			Where: "NewRegisteredSucceededEvent",
			What:  "failed to create event message",
			Why: errs.Meta{
				"Routing Key": RegisteredSucceededEventTypeRoutingKey,
			},
			Who: err,
		})
	}

	return message.NewMessage(RegisteredSucceededEventTypeRoutingKey, attributesJson, []byte{}), nil
}
