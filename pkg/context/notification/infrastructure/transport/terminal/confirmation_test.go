package terminal_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/notification/infrastructure/transport/terminal"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events/user"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/transfers"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/records"
)

type ConfirmationTestSuite struct {
	suite.Suite
	sut          transfers.Transfer[*user.CreatedSucceededAttributes]
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
	attributes := new(user.CreatedSucceededAttributes)

	messages.RandomAttributes(attributes)

	link := fmt.Sprintf("Hi %s, please confirm your account through this link: %s/v4/account/verify/%s", attributes.Username, suite.appServerURL, attributes.ID)

	suite.logger.Mock.On("Info", link)

	suite.NoError(suite.sut.Submit(attributes))

	suite.logger.AssertExpectations(suite.T())
}

func TestIntegrationConfirmationSuite(t *testing.T) {
	suite.Run(t, new(ConfirmationTestSuite))
}
