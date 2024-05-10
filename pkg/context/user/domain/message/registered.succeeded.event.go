package message

import (
	"encoding/json"

	"github.com/bastean/codexgo/pkg/context/shared/domain/serror"
	"github.com/bastean/codexgo/pkg/context/shared/domain/smessage"
)

var RegisteredSucceededEventTypeRoutingKey = smessage.NewRoutingKey(&smessage.MessageRoutingKey{
	Module:    "user",
	Version:   "1",
	Type:      smessage.Type.Event,
	Aggregate: "user",
	Event:     "registered",
	Status:    smessage.Status.Succeeded,
})

type RegisteredSucceededEventAttributes struct {
	Id       string
	Email    string
	Username string
}

func NewRegisteredSucceededEvent(attributes *RegisteredSucceededEventAttributes) (*smessage.Message, error) {
	attributesJson, err := json.Marshal(attributes)

	if err != nil {
		return nil, serror.NewInternal(&serror.Bubble{
			Where: "NewRegisteredSucceededEvent",
			What:  "failure to create an event message",
			Why: serror.Meta{
				"Routing Key": RegisteredSucceededEventTypeRoutingKey,
			},
			Who: err,
		})
	}

	return smessage.NewMessage(RegisteredSucceededEventTypeRoutingKey, attributesJson, []byte{}), nil
}
