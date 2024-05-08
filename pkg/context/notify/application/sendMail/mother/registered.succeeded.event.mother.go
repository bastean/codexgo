package eventMother

import (
	"encoding/json"

	"github.com/bastean/codexgo/pkg/context/notify/application/sendMail"
	"github.com/bastean/codexgo/pkg/context/shared/domain/message"
	"github.com/bastean/codexgo/pkg/context/shared/domain/service/mother"
)

var RegisteredSucceededEventTypeRoutingKey = message.NewRoutingKey(&message.MessageRoutingKey{
	Module:    "user",
	Version:   "1",
	Type:      message.Type.Event,
	Aggregate: "user",
	Event:     "registered",
	Status:    message.Status.Succeeded,
})

func Random() *message.Message {
	id := mother.Create.UUID()
	email := mother.Create.Email()
	username := mother.Create.Username()

	attributes := sendMail.RegisteredSucceededEventAttributes{
		Id:       id,
		Email:    email,
		Username: username,
	}

	attributesJson, _ := json.Marshal(attributes)

	return message.NewMessage(RegisteredSucceededEventTypeRoutingKey, attributesJson, []byte{})
}
