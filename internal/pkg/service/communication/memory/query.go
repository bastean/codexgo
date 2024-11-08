package memory

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/queries"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications/memory"
)

type (
	QueryBus     = memory.QueryBus
	QueryHandler = queries.Handler
)

var (
	NewQueryBus = memory.NewQueryBus
)
