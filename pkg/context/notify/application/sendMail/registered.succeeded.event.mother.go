package sendMail

import (
	"encoding/json"

	"github.com/bastean/codexgo/pkg/context/shared/domain/smessage"
	"github.com/bastean/codexgo/pkg/context/shared/domain/sservice"
)

var RegisteredSucceededEventTypeRoutingKey = smessage.NewRoutingKey(&smessage.MessageRoutingKey{
	Module:    "user",
	Version:   "1",
	Type:      smessage.Type.Event,
	Aggregate: "user",
	Event:     "registered",
	Status:    smessage.Status.Succeeded,
})

func RandomEvent() *smessage.Message {
	id := sservice.Create.UUID()
	email := sservice.Create.Email()
	username := sservice.Create.Username()

	attributes := RegisteredSucceededEventAttributes{
		Id:       id,
		Email:    email,
		Username: username,
	}

	attributesJson, _ := json.Marshal(attributes)

	return smessage.NewMessage(RegisteredSucceededEventTypeRoutingKey, attributesJson, []byte{})
}
