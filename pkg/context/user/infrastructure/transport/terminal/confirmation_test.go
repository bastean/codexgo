package terminal_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/transfers"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/records"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/transport/terminal"
)

type ConfirmationTestSuite struct {
	suite.Suite
	sut          transfers.Transfer
	logger       *records.LoggerMock
	appServerURL string
}

func (suite *ConfirmationTestSuite) SetupTest() {
	suite.logger = new(records.LoggerMock)

	suite.appServerURL = os.Getenv("CODEXGO_SERVER_GIN_URL")

	suite.sut = &terminal.Confirmation{
		Logger:       suite.logger,
		AppServerURL: suite.appServerURL,
	}
}

func (suite *ConfirmationTestSuite) TestSubmit() {
	message := user.RandomCreatedSucceeded()

	event := new(user.CreatedSucceeded)

	event.Attributes = new(user.CreatedSucceededAttributes)

	suite.NoError(json.Unmarshal(message.Attributes, event.Attributes))

	link := fmt.Sprintf("Hi %s, please confirm your account through this link: %s/v4/account/verify/%s", event.Attributes.Username, suite.appServerURL, event.Attributes.Id)

	suite.logger.Mock.On("Info", link)

	suite.NoError(suite.sut.Submit(event.Attributes))

	suite.logger.AssertExpectations(suite.T())
}

func TestIntegrationConfirmationSuite(t *testing.T) {
	suite.Run(t, new(ConfirmationTestSuite))
}
