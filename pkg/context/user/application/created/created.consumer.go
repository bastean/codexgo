package created

import (
	"encoding/json"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/domain/types"
	"github.com/bastean/codexgo/pkg/context/user/domain/event"
)

type Consumer struct {
	models.UseCase[*event.CreatedSucceeded, types.Empty]
	Queues []*messages.Queue
}

func (consumer *Consumer) SubscribedTo() []*messages.Queue {
	return consumer.Queues
}

func (consumer *Consumer) On(message *messages.Message) error {
	user := new(event.CreatedSucceeded)

	user.Attributes = new(event.CreatedSucceededAttributes)

	err := json.Unmarshal(message.Attributes, user.Attributes)

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "On",
			What:  "failure to obtain message attributes",
			Why: errors.Meta{
				"Id":          message.Id,
				"Routing Key": message.Type,
				"Occurred On": message.OccurredOn,
			},
			Who: err,
		})
	}

	_, err = consumer.UseCase.Run(user)

	if err != nil {
		return errors.BubbleUp(err, "On")
	}

	return nil
}
