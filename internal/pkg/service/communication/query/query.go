package query

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications/memory"
)

var (
	Bus roles.QueryBus
)

var (
	NewBus = memory.NewQueryBus
)

type (
	Mapper  = memory.QueryMapper
	Handler = roles.QueryHandler
)

func New(key messages.Key, attributes, meta any) *messages.Message {
	return messages.New(key, attributes, meta)
}
