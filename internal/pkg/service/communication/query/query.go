package query

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/queries"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications/memory"
)

var (
	Bus queries.Bus
)

var (
	NewBus = memory.NewQueryBus
)

type (
	Mapper  = memory.QueryMapper
	Handler = queries.Handler
)

func New(key queries.Key, attributes, meta any) *queries.Query {
	return messages.New[queries.Query](key, attributes, meta)
}
