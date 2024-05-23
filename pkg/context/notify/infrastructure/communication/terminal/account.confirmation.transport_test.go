package terminal_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/bastean/codexgo/pkg/context/notify/application/send"
	"github.com/bastean/codexgo/pkg/context/notify/domain/model"
	"github.com/bastean/codexgo/pkg/context/notify/infrastructure/communication/terminal"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/slogger"
	"github.com/stretchr/testify/suite"
)

type NotifyTerminalAccountConfirmationTransportTestSuite struct {
	suite.Suite
	sut       model.Transport
	logger    *slogger.LoggerMock
	serverURL string
}

func (suite *NotifyTerminalAccountConfirmationTransportTestSuite) SetupTest() {
	suite.serverURL = os.Getenv("URL")

	suite.logger = new(slogger.LoggerMock)

	suite.sut = &terminal.AccountConfirmation{
		ServerURL: suite.serverURL,
		Logger:    suite.logger,
	}
}

func (suite *NotifyTerminalAccountConfirmationTransportTestSuite) TestSubmit() {
	message := send.RandomEvent()

	attributes := new(send.CreatedSucceededEventAttributes)

	json.Unmarshal(message.Attributes, attributes)

	confirmationLink := fmt.Sprintf("Hi %s, please confirm your account using this link: %s/verify/%s", attributes.Username, suite.serverURL, attributes.Id)

	suite.logger.Mock.On("Info", confirmationLink)

	suite.NoError(suite.sut.Submit(attributes))

	suite.logger.AssertExpectations(suite.T())
}

func TestIntegrationNotifyTerminalAccountConfirmationTransportSuite(t *testing.T) {
	suite.Run(t, new(NotifyTerminalAccountConfirmationTransportTestSuite))
}
