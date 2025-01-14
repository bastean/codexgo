package roles

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
)

type EventConsumer interface {
	On(*messages.Message) error
}

type EventBus interface {
	Subscribe(messages.Key, EventConsumer) error
	Publish(*messages.Message) error
}
