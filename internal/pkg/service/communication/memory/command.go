package memory

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/command"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications/memory"
)

type (
	CommandBus     = memory.CommandBus
	CommandHandler = command.Handler
)

var (
	NewCommandBus = memory.NewCommandBus
)
