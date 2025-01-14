package roles

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
)

type CommandHandler interface {
	Handle(*messages.Message) error
}

type CommandBus interface {
	Register(messages.Key, CommandHandler) error
	Dispatch(*messages.Message) error
}
