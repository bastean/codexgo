package sendMail

import (
	"encoding/json"

	"github.com/bastean/codexgo/pkg/context/notify/domain/model"
	"github.com/bastean/codexgo/pkg/context/notify/domain/template"
	"github.com/bastean/codexgo/pkg/context/shared/domain/errs"
	"github.com/bastean/codexgo/pkg/context/shared/domain/message"
	sharedModel "github.com/bastean/codexgo/pkg/context/shared/domain/model"
	"github.com/bastean/codexgo/pkg/context/shared/domain/queue"
	"github.com/bastean/codexgo/pkg/context/shared/domain/types"
)

type RegisteredSucceededEventConsumer struct {
	sharedModel.UseCase[model.MailTemplate, *types.Empty]
	Queues []*queue.Queue
}

func (consumer *RegisteredSucceededEventConsumer) SubscribedTo() []*queue.Queue {
	return consumer.Queues
}

func (consumer *RegisteredSucceededEventConsumer) On(message *message.Message) error {
	attributes := new(RegisteredSucceededEventAttributes)

	err := json.Unmarshal(message.Attributes, attributes)

	if err != nil {
		return errs.NewFailedError(&errs.Bubble{
			Where: "On",
			What:  "failed getting message attributes",
			Why: errs.Meta{
				"Message": message.Id,
			},
			Who: err,
		})
	}

	accountConfirmationTemplate := &template.AccountConfirmationMail{
		Mail: &template.Mail{
			To: []string{attributes.Email},
		},
		Username:         attributes.Username,
		ConfirmationLink: attributes.Id,
	}

	_, err = consumer.UseCase.Run(accountConfirmationTemplate)

	if err != nil {
		return errs.BubbleUp(err, "On")
	}

	return nil
}
