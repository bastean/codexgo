package user

import (
	"github.com/bastean/codexgo/v4/internal/pkg/service/communication/rabbitmq"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
)

var QueueSendConfirmation = &messages.Queue{
	Name: messages.NewRecipientName(&messages.RecipientNameComponents{
		Service: "user",
		Entity:  "user",
		Action:  "send confirmation",
		Event:   "created",
		Status:  "succeeded",
	}),
	Bindings: []string{rabbitmq.BindingEventCreatedSucceeded},
}
