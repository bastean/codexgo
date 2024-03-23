package message

import (
	"encoding/json"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/message"
)

var RegisteredSucceededEventRoutingKey = message.NewMessageRoutingKey(&message.MessageRoutingKey{Module: "user", Version: "1", Type: message.Event, Aggregate: "user", Event: "registered", Status: message.Succeeded})

var FailedRegisteredSucceededEvent = errors.Failed{Message: "Failed message creation for " + RegisteredSucceededEventRoutingKey}

type RegisteredSucceededEventAttributes struct {
	Id       string
	Email    string
	Username string
}

func NewRegisteredSucceededEvent(attributes *RegisteredSucceededEventAttributes) *message.Message {
	attributesJson, err := json.Marshal(attributes)

	if err != nil {
		panic(FailedRegisteredSucceededEvent)
	}

	return message.NewMessage(RegisteredSucceededEventRoutingKey, attributesJson, []byte{})
}
