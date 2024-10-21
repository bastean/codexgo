package memory

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/query"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications/memory"
)

type (
	QueryBus     = memory.QueryBus
	QueryHandler = query.Handler
)

var (
	NewQueryBus = memory.NewQueryBus
)
