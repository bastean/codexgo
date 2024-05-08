package sendMail_test

import (
	"encoding/json"
	"testing"

	"github.com/bastean/codexgo/pkg/context/notify/application/sendMail"
	eventMother "github.com/bastean/codexgo/pkg/context/notify/application/sendMail/mother"
	"github.com/bastean/codexgo/pkg/context/notify/domain/template"
	communicationMock "github.com/bastean/codexgo/pkg/context/notify/infrastructure/communication/mock"
	"github.com/bastean/codexgo/pkg/context/shared/domain/model"
	"github.com/bastean/codexgo/pkg/context/shared/domain/queue"
	"github.com/stretchr/testify/suite"
)

type RegisteredSucceededEventConsumerTestSuite struct {
	suite.Suite
	sut      model.Consumer
	sendMail *sendMail.SendMail
	mail     *communicationMock.MailMock
	queues   []*queue.Queue
}

func (suite *RegisteredSucceededEventConsumerTestSuite) SetupTest() {
	queueName := queue.NewQueueName(&queue.QueueName{
		Module: "queue",
		Action: "assert",
		Event:  "test.succeeded",
	})
	suite.queues = append(suite.queues, &queue.Queue{Name: queueName})
	suite.mail = new(communicationMock.MailMock)
	suite.sendMail = &sendMail.SendMail{Mail: suite.mail}
	suite.sut = &sendMail.RegisteredSucceededEventConsumer{
		UseCase: suite.sendMail,
		Queues:  suite.queues,
	}
}

func (suite *RegisteredSucceededEventConsumerTestSuite) TestEventConsumer() {
	message := eventMother.Random()

	attributes := new(sendMail.RegisteredSucceededEventAttributes)

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

func TestUnitRegisteredSucceededEventConsumerSuite(t *testing.T) {
	suite.Run(t, new(RegisteredSucceededEventConsumerTestSuite))
}
