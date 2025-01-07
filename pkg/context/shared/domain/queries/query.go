package queries

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
)

type Handler interface {
	Handle(*messages.Message) (*messages.Message, error)
}

type Bus interface {
	Register(messages.Key, Handler) error
	Ask(*messages.Message) (*messages.Message, error)
}
