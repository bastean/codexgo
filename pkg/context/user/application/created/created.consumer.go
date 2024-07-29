package created

import (
	"encoding/json"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/usecase"
)

type Consumer struct {
	usecase.Created
	Queues []*messages.Queue
}

func (consumer *Consumer) SubscribedTo() []*messages.Queue {
	return consumer.Queues
}

func (consumer *Consumer) On(message *messages.Message) error {
	event := new(user.CreatedSucceeded)

	event.Attributes = new(user.CreatedSucceededAttributes)

	err := json.Unmarshal(message.Attributes, event.Attributes)

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "On",
			What:  "Failure to obtain message attributes",
			Why: errors.Meta{
				"Id":          message.Id,
				"Routing Key": message.Type,
				"Occurred On": message.OccurredOn,
			},
			Who: err,
		})
	}

	err = consumer.Created.Run(event)

	if err != nil {
		return errors.BubbleUp(err, "On")
	}

	return nil
}
