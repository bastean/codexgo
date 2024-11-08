package user

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
)

var CreatedSucceededKey = messages.NewKey(&messages.KeyComponents{
	Service: "user",
	Version: "1",
	Type:    messages.Type.Event,
	Entity:  "user",
	Event:   "created",
	Status:  messages.Status.Succeeded,
})

type CreatedSucceededAttributes struct {
	ID, Email, Username string
}

type CreatedSucceededMeta struct{}
