package user

import (
	"encoding/json"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
)

var CreatedSucceededTypeRoutingKey = messages.NewRoutingKey(&messages.RoutingKeyComponents{
	Service: "user",
	Version: "1",
	Type:    messages.Type.Event,
	Entity:  "user",
	Event:   "created",
	Status:  messages.Status.Succeeded,
})

type CreatedSucceededAttributes struct {
	Id, Email, Username string
}

type CreatedSucceeded struct {
	Attributes *CreatedSucceededAttributes
}

func NewCreatedSucceeded(event *CreatedSucceeded) (*messages.Message, error) {
	attributes, err := json.Marshal(event.Attributes)

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "NewCreatedSucceeded",
			What:  "Failure to create event message attributes",
			Why: errors.Meta{
				"Routing Key": CreatedSucceededTypeRoutingKey,
			},
			Who: err,
		})
	}

	return messages.New(CreatedSucceededTypeRoutingKey, attributes, messages.Meta{}), nil
}
