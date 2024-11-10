package commands

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
)

type (
	Key = messages.Key
)

type Command messages.Message

type Handler interface {
	Handle(*Command) error
}

type Bus interface {
	Register(Key, Handler) error
	Dispatch(*Command) error
}
