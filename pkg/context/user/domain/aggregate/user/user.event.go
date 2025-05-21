package user

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

var CreatedSucceededKey, _ = values.New[*messages.Key](messages.ParseKey(&messages.KeyComponents{
	Service: "user",
	Version: "1",
	Type:    messages.Type.Event,
	Entity:  "user",
	Action:  "created",
	Status:  messages.Status.Succeeded,
}))

type CreatedSucceededAttributes = struct {
	VerifyToken, ID, Email, Username string
}

type CreatedSucceededMeta = struct{}

var ResetQueuedKey, _ = values.New[*messages.Key](messages.ParseKey(&messages.KeyComponents{
	Service: "user",
	Version: "1",
	Type:    messages.Type.Event,
	Entity:  "user",
	Action:  "reset",
	Status:  messages.Status.Queued,
}))

type ResetQueuedAttributes = struct {
	ResetToken, ID, Email, Username string
}

type ResetQueuedMeta = struct{}
