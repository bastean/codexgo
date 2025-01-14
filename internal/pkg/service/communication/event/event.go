package event

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications/memory"
)

var (
	NewBus = memory.NewEventBus
)

type (
	Consumer = roles.EventConsumer
	Bus      = roles.EventBus
	Mapper   = memory.EventMapper
)
