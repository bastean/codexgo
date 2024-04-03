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
	queue    []*queue.Queue
}

func (suite *RegisteredSucceededEventConsumerTestSuite) SetupTest() {
	queueName := queue.NewQueueName(&queue.QueueName{Module: "queue", Action: "assert", Event: "test.succeeded"})
	suite.queue = append(suite.queue, queue.NewQueue(queueName))
	suite.mail = communicationMock.NewMailMock()
	suite.sendMail = sendMail.NewSendMail(suite.mail)
	suite.sut = sendMail.NewRegisteredSucceededEventConsumer(suite.sendMail, suite.queue)
}

func (suite *RegisteredSucceededEventConsumerTestSuite) TestEventConsumer() {
	message := eventMother.Random()

	attributes := new(sendMail.RegisteredSucceededEventAttributes)

	json.Unmarshal(message.Attributes, attributes)

	mailTemplate := template.NewMail([]string{attributes.Email})

	accountConfirmationTemplate := template.NewAccountConfirmationMail(mailTemplate, attributes.Username, attributes.Id)

	suite.mail.On("Send", accountConfirmationTemplate)

	suite.sut.On(message)

	suite.mail.AssertExpectations(suite.T())
}

func TestUnitRegisteredSucceededEventConsumerSuite(t *testing.T) {
	suite.Run(t, new(RegisteredSucceededEventConsumerTestSuite))
}
