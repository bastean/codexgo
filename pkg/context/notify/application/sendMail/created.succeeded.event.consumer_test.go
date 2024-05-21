package sendMail_test

import (
	"encoding/json"
	"testing"

	"github.com/bastean/codexgo/pkg/context/notify/application/sendMail"
	"github.com/bastean/codexgo/pkg/context/notify/domain/model"
	"github.com/bastean/codexgo/pkg/context/notify/domain/template"
	"github.com/bastean/codexgo/pkg/context/notify/infrastructure/communication"
	"github.com/bastean/codexgo/pkg/context/shared/domain/smodel"
	"github.com/bastean/codexgo/pkg/context/shared/domain/squeue"
	"github.com/bastean/codexgo/pkg/context/shared/domain/stype"
	"github.com/stretchr/testify/suite"
)

type CreatedSucceededEventConsumerTestSuite struct {
	suite.Suite
	sut     smodel.Consumer
	useCase smodel.UseCase[model.MailTemplate, *stype.Empty]
	mail    *communication.MailMock
	queues  []*squeue.Queue
}

func (suite *CreatedSucceededEventConsumerTestSuite) SetupTest() {
	queueName := squeue.NewQueueName(&squeue.QueueName{
		Module: "queue",
		Action: "assert",
		Event:  "test.succeeded",
	})
	suite.queues = append(suite.queues, &squeue.Queue{
		Name: queueName,
	})
	suite.mail = new(communication.MailMock)
	suite.useCase = &sendMail.SendMail{
		Mail: suite.mail,
	}
	suite.sut = &sendMail.CreatedSucceededEventConsumer{
		UseCase: suite.useCase,
		Queues:  suite.queues,
	}
}

func (suite *CreatedSucceededEventConsumerTestSuite) TestEventConsumer() {
	message := sendMail.RandomEvent()

	attributes := new(sendMail.CreatedSucceededEventAttributes)

	json.Unmarshal(message.Attributes, attributes)

	accountConfirmationTemplate := &template.AccountConfirmationMail{
		Mail: &template.Mail{
			To: []string{attributes.Email},
		},
		Username:         attributes.Username,
		ConfirmationLink: attributes.Id,
	}

	suite.mail.On("Send", accountConfirmationTemplate)

	suite.NoError(suite.sut.On(message))

	suite.mail.AssertExpectations(suite.T())
}

func TestUnitCreatedSucceededEventConsumerSuite(t *testing.T) {
	suite.Run(t, new(CreatedSucceededEventConsumerTestSuite))
}
