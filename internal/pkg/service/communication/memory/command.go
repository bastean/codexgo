package memory

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/commands"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications/memory"
)

type (
	CommandBus     = memory.CommandBus
	CommandHandler = commands.Handler
)

var (
	NewCommandBus = memory.NewCommandBus
)
