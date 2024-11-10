package events

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
)

type (
	Key       = messages.Key
	Recipient = messages.Recipient
)

type Event messages.Message

type Consumer interface {
	On(*Event) error
}

type Bus interface {
	Subscribe(Key, Consumer) error
	Publish(*Event) error
}
