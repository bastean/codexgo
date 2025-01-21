package events

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
)

var UserCreatedSucceededKey = messages.NewKey(&messages.KeyComponents{
	Service: "user",
	Version: "1",
	Type:    messages.Type.Event,
	Entity:  "user",
	Event:   "created",
	Status:  messages.Status.Succeeded,
})

type (
	UserCreatedSucceededAttributes struct {
		Verify, ID, Email, Username string
	}
	UserCreatedSucceededMeta struct{}
)

var UserResetQueuedKey = messages.NewKey(&messages.KeyComponents{
	Service: "user",
	Version: "1",
	Type:    messages.Type.Event,
	Entity:  "user",
	Event:   "reset",
	Status:  messages.Status.Queued,
})

type (
	UserResetQueuedAttributes struct {
		Reset, ID, Email, Username string
	}
	UserResetQueuedMeta struct{}
)
