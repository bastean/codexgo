package command

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications/memory"
)

var (
	Bus roles.CommandBus
)

var (
	NewBus = memory.NewCommandBus
)

type (
	Mapper  = memory.CommandMapper
	Handler = roles.CommandHandler
)

func New(key messages.Key, attributes, meta any) *messages.Message {
	return messages.New(key, attributes, meta)
}
