package queries

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
)

type (
	Key = messages.Key
)

type (
	Query    messages.Message
	Response messages.Message
)

type Handler interface {
	Handle(*Query) (*Response, error)
}

type Bus interface {
	Register(Key, Handler) error
	Ask(*Query) (*Response, error)
}
