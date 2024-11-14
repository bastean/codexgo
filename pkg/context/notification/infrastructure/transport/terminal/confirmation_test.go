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

func (s *ConfirmationTestSuite) SetupTest() {
	s.logger = new(records.LoggerMock)

	s.appServerURL = os.Getenv("CODEXGO_SERVER_GIN_URL")

	s.sut = &terminal.Confirmation{
		Logger:       s.logger,
		AppServerURL: s.appServerURL,
	}
}

func (s *ConfirmationTestSuite) TestSubmit() {
	attributes := new(user.CreatedSucceededAttributes)

	messages.RandomAttributes(attributes)

	link := fmt.Sprintf("Hi %s, please confirm your account through this link: %s/v4/account/verify/%s", attributes.Username, s.appServerURL, attributes.ID)

	s.logger.Mock.On("Info", link)

	s.NoError(s.sut.Submit(attributes))

	s.logger.AssertExpectations(s.T())
}

func TestIntegrationConfirmationSuite(t *testing.T) {
	suite.Run(t, new(ConfirmationTestSuite))
}
