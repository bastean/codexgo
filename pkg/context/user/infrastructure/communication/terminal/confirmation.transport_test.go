package terminal_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/loggers"
	"github.com/bastean/codexgo/pkg/context/user/domain/event"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/communication/terminal"
	"github.com/stretchr/testify/suite"
)

type TerminalConfirmationTransportTestSuite struct {
	suite.Suite
	sut       models.Transport
	logger    *loggers.LoggerMock
	serverURL string
}

func (suite *TerminalConfirmationTransportTestSuite) SetupTest() {
	suite.logger = new(loggers.LoggerMock)

	suite.serverURL = os.Getenv("URL")

	suite.sut = &terminal.Confirmation{
		Logger:    suite.logger,
		ServerURL: suite.serverURL,
	}
}

func (suite *TerminalConfirmationTransportTestSuite) TestSubmit() {
	message := event.RandomCreatedSucceeded()

	user := new(event.CreatedSucceeded)

	user.Attributes = new(event.CreatedSucceededAttributes)

	suite.NoError(json.Unmarshal(message.Attributes, user.Attributes))

	link := fmt.Sprintf("Hi %s, please confirm your account through this link: %s/verify/%s", user.Attributes.Username, suite.serverURL, user.Attributes.Id)

	suite.logger.Mock.On("Info", link)

	suite.NoError(suite.sut.Submit(user.Attributes))

	suite.logger.AssertExpectations(suite.T())
}

func TestIntegrationTerminalConfirmationTransportSuite(t *testing.T) {
	suite.Run(t, new(TerminalConfirmationTransportTestSuite))
}
