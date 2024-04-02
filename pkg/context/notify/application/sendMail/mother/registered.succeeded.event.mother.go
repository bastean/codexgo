package eventMother

import (
	"encoding/json"

	"github.com/bastean/codexgo/pkg/context/notify/application/sendMail"
	"github.com/bastean/codexgo/pkg/context/shared/domain/message"
	"github.com/bastean/codexgo/pkg/context/shared/domain/service/mother"
)

var RegisteredSucceededEventRoutingKey = message.NewMessageRoutingKey(&message.MessageRoutingKey{Module: "user", Version: "1", Type: message.Event, Aggregate: "user", Event: "registered", Status: message.Succeeded})

func Random() *message.Message {
	id := mother.Create.UUID()
	email := mother.Create.Email()
	username := mother.Create.Username()

	attributes := sendMail.NewRegisteredSucceededEventAttributes(id, email, username)

	attributesJson, err := json.Marshal(attributes)

	if err != nil {
		panic(err)
	}

	return message.NewMessage(RegisteredSucceededEventRoutingKey, attributesJson, []byte{})
}
