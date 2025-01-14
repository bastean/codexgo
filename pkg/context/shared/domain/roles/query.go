package roles

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
)

type QueryHandler interface {
	Handle(*messages.Message) (*messages.Message, error)
}

type QueryBus interface {
	Register(messages.Key, QueryHandler) error
	Ask(*messages.Message) (*messages.Message, error)
}
