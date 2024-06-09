package send

import (
	"encoding/json"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/domain/queues"
	"github.com/bastean/codexgo/pkg/context/shared/domain/types"
)

type CreatedSucceededEventConsumer struct {
	models.UseCase[any, types.Empty]
	Queues []*queues.Queue
}

func (consumer *CreatedSucceededEventConsumer) SubscribedTo() []*queues.Queue {
	return consumer.Queues
}

func (consumer *CreatedSucceededEventConsumer) On(message *messages.Message) error {
	attributes := new(CreatedSucceededEventAttributes)

	err := json.Unmarshal(message.Attributes, attributes)

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "On",
			What:  "failure to obtain message attributes",
			Why: errors.Meta{
				"Message": message.Id,
			},
			Who: err,
		})
	}

	_, err = consumer.UseCase.Run(attributes)

	if err != nil {
		return errors.BubbleUp(err, "On")
	}

	return nil
}
