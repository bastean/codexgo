package message

import (
	"encoding/json"

	"github.com/bastean/codexgo/pkg/context/shared/domain/serror"
	"github.com/bastean/codexgo/pkg/context/shared/domain/smessage"
)

var CreatedSucceededEventTypeRoutingKey = smessage.NewRoutingKey(&smessage.MessageRoutingKey{
	Module:    "user",
	Version:   "1",
	Type:      smessage.Type.Event,
	Aggregate: "user",
	Event:     "created",
	Status:    smessage.Status.Succeeded,
})

type CreatedSucceededEventAttributes struct {
	Id       string
	Email    string
	Username string
}

func NewCreatedSucceededEvent(attributes *CreatedSucceededEventAttributes) (*smessage.Message, error) {
	attributesJson, err := json.Marshal(attributes)

	if err != nil {
		return nil, serror.NewInternal(&serror.Bubble{
			Where: "NewCreatedSucceededEvent",
			What:  "failure to create an event message",
			Why: serror.Meta{
				"Routing Key": CreatedSucceededEventTypeRoutingKey,
			},
			Who: err,
		})
	}

	return smessage.NewMessage(CreatedSucceededEventTypeRoutingKey, attributesJson, []byte{}), nil
}
