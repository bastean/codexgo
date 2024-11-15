package event

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications/memory"
)

var (
	NewBus = memory.NewEventBus
)

type (
	Consumer = events.Consumer
	Bus      = events.Bus
	Mapper   = memory.EventMapper
)
