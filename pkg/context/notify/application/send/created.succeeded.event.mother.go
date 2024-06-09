package send

import (
	"encoding/json"

	"github.com/bastean/codexgo/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/pkg/context/shared/domain/services"
)

var CreatedSucceededEventTypeRoutingKey = messages.NewRoutingKey(&messages.MessageRoutingKey{
	Module:    "user",
	Version:   "1",
	Type:      messages.Type.Event,
	Aggregate: "user",
	Event:     "created",
	Status:    messages.Status.Succeeded,
})

func RandomEvent() *messages.Message {
	id := services.Create.UUID()
	email := services.Create.Email()
	username := services.Create.Username()

	attributes := CreatedSucceededEventAttributes{
		Id:       id,
		Email:    email,
		Username: username,
	}

	attributesJson, _ := json.Marshal(attributes)

	return messages.NewMessage(CreatedSucceededEventTypeRoutingKey, attributesJson, messages.Meta{})
}
