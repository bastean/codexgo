package terminal_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/bastean/codexgo/pkg/context/notify/application/send"
	"github.com/bastean/codexgo/pkg/context/notify/domain/model"
	"github.com/bastean/codexgo/pkg/context/notify/infrastructure/communication/terminal"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/loggers"
	"github.com/stretchr/testify/suite"
)

type NotifyTerminalAccountConfirmationTransportTestSuite struct {
	suite.Suite
	sut       model.Transport
	logger    *loggers.LoggerMock
	serverURL string
}

func (suite *NotifyTerminalAccountConfirmationTransportTestSuite) SetupTest() {
	suite.logger = new(loggers.LoggerMock)

	suite.serverURL = os.Getenv("URL")

	suite.sut = &terminal.AccountConfirmation{
		Logger:    suite.logger,
		ServerURL: suite.serverURL,
	}
}

func (suite *NotifyTerminalAccountConfirmationTransportTestSuite) TestSubmit() {
	message := send.RandomEvent()

	attributes := new(send.CreatedSucceededEventAttributes)

	json.Unmarshal(message.Attributes, attributes)

	confirmationLink := fmt.Sprintf("Hi %s, please confirm your account through this link: %s/verify/%s", attributes.Username, suite.serverURL, attributes.Id)

	suite.logger.Mock.On("Info", confirmationLink)

	suite.NoError(suite.sut.Submit(attributes))

	suite.logger.AssertExpectations(suite.T())
}

func TestIntegrationNotifyTerminalAccountConfirmationTransportSuite(t *testing.T) {
	suite.Run(t, new(NotifyTerminalAccountConfirmationTransportTestSuite))
}
