package sendMail

import (
	"encoding/json"

	"github.com/bastean/codexgo/pkg/context/notify/domain/model"
	"github.com/bastean/codexgo/pkg/context/notify/domain/template"
	"github.com/bastean/codexgo/pkg/context/shared/domain/serror"
	"github.com/bastean/codexgo/pkg/context/shared/domain/smessage"
	"github.com/bastean/codexgo/pkg/context/shared/domain/smodel"
	"github.com/bastean/codexgo/pkg/context/shared/domain/squeue"
	"github.com/bastean/codexgo/pkg/context/shared/domain/stype"
)

type RegisteredSucceededEventConsumer struct {
	smodel.UseCase[model.MailTemplate, *stype.Empty]
	Queues []*squeue.Queue
}

func (consumer *RegisteredSucceededEventConsumer) SubscribedTo() []*squeue.Queue {
	return consumer.Queues
}

func (consumer *RegisteredSucceededEventConsumer) On(message *smessage.Message) error {
	attributes := new(RegisteredSucceededEventAttributes)

	err := json.Unmarshal(message.Attributes, attributes)

	if err != nil {
		return serror.NewFailedError(&serror.Bubble{
			Where: "On",
			What:  "failed getting message attributes",
			Why: serror.Meta{
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
		return serror.BubbleUp(err, "On")
	}

	return nil
}
