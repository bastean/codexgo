package sendMail

import (
	"encoding/json"

	"github.com/bastean/codexgo/pkg/context/notify/domain/template"
	"github.com/bastean/codexgo/pkg/context/shared/domain/message"
	"github.com/bastean/codexgo/pkg/context/shared/domain/model"
	"github.com/bastean/codexgo/pkg/context/shared/domain/queue"
)

type RegisteredSucceededEventConsumer struct {
	*SendMail
	Queues []*queue.Queue
}

func (consumer *RegisteredSucceededEventConsumer) SubscribedTo() []*queue.Queue {
	return consumer.Queues
}

func (consumer *RegisteredSucceededEventConsumer) On(message *message.Message) {
	attributes := new(RegisteredSucceededEventAttributes)

	json.Unmarshal(message.Attributes, attributes)

	mailTemplate := template.NewMail([]string{attributes.Email})
	accountConfirmationTemplate := template.NewAccountConfirmationMail(mailTemplate, attributes.Username, attributes.Id)

	consumer.SendMail.Run(accountConfirmationTemplate)
}

func NewRegisteredSucceededEventConsumer(sendEmail *SendMail, queues []*queue.Queue) model.Consumer {
	return &RegisteredSucceededEventConsumer{
		SendMail: sendEmail,
		Queues:   queues,
	}
}
