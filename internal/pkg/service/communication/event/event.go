package event

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications/memory"
)

var (
	Bus events.Bus
)

var (
	NewBus = memory.NewEventBus
)

type (
	Mapper   = memory.EventMapper
	Consumer = events.Consumer
)
