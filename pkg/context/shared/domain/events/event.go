package events

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
)

type Consumer interface {
	On(*messages.Message) error
}

type Bus interface {
	Subscribe(messages.Key, Consumer) error
	Publish(*messages.Message) error
}
