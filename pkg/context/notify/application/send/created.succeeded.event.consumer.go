package send

import (
	"encoding/json"

	"github.com/bastean/codexgo/pkg/context/shared/domain/serror"
	"github.com/bastean/codexgo/pkg/context/shared/domain/smessage"
	"github.com/bastean/codexgo/pkg/context/shared/domain/smodel"
	"github.com/bastean/codexgo/pkg/context/shared/domain/squeue"
	"github.com/bastean/codexgo/pkg/context/shared/domain/stype"
)

type CreatedSucceededEventConsumer struct {
	smodel.UseCase[any, *stype.Empty]
	Queues []*squeue.Queue
}

func (consumer *CreatedSucceededEventConsumer) SubscribedTo() []*squeue.Queue {
	return consumer.Queues
}

func (consumer *CreatedSucceededEventConsumer) On(message *smessage.Message) error {
	attributes := new(CreatedSucceededEventAttributes)

	err := json.Unmarshal(message.Attributes, attributes)

	if err != nil {
		return serror.NewInternal(&serror.Bubble{
			Where: "On",
			What:  "failure to obtain message attributes",
			Why: serror.Meta{
				"Message": message.Id,
			},
			Who: err,
		})
	}

	_, err = consumer.UseCase.Run(attributes)

	if err != nil {
		return serror.BubbleUp(err, "On")
	}

	return nil
}
