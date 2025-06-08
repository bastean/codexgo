package memory

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/commands"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
)

type CommandBus struct {
	Handlers commands.Mapper
}

func (b *CommandBus) Register(key *messages.Key, handler roles.CommandHandler) error {
	_, ok := b.Handlers[key.Value()]

	if ok {
		return errors.New[errors.Internal](&errors.Bubble{
			What: "Already registered",
			Why: errors.Meta{
				"Key": key.Value(),
			},
		})
	}

	b.Handlers[key.Value()] = handler

	return nil
}

func (b *CommandBus) Dispatch(command *messages.Message) error {
	handler, ok := b.Handlers[command.Key.Value()]

	if !ok {
		return errors.New[errors.Internal](&errors.Bubble{
			What: "Failure to execute a Command without a Handler",
			Why: errors.Meta{
				"ID":  command.ID.Value(),
				"Key": command.Key.Value(),
			},
		})
	}

	err := handler.Handle(command)

	if err != nil {
		return errors.BubbleUp(err)
	}

	return nil
}
