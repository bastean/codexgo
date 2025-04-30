package events

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

var UserCreatedSucceededKey, _ = values.New[*messages.Key](messages.ParseKey(&messages.KeyComponents{
	Service: "user",
	Version: "1",
	Type:    messages.Type.Event,
	Entity:  "user",
	Action:  "created",
	Status:  messages.Status.Succeeded,
}))

type (
	UserCreatedSucceededAttributes struct {
		Verify, ID, Email, Username string
	}
	UserCreatedSucceededMeta struct{}
)

var UserResetQueuedKey, _ = values.New[*messages.Key](messages.ParseKey(&messages.KeyComponents{
	Service: "user",
	Version: "1",
	Type:    messages.Type.Event,
	Entity:  "user",
	Action:  "reset",
	Status:  messages.Status.Queued,
}))

type (
	UserResetQueuedAttributes struct {
		Reset, ID, Email, Username string
	}
	UserResetQueuedMeta struct{}
)
