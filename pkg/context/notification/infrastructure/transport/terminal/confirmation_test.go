package terminal_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/notification/domain/aggregate/recipient"
	"github.com/bastean/codexgo/v4/pkg/context/notification/infrastructure/transport"
	"github.com/bastean/codexgo/v4/pkg/context/notification/infrastructure/transport/terminal"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/records"
)

type ConfirmationTestSuite struct {
	transport.OfflineSuite
}

func (s *ConfirmationTestSuite) SetupSuite() {
	appServerURL := os.Getenv("CODEXGO_SERVER_GIN_URL")

	s.Recipient = recipient.Mother().RecipientValid()

	s.Recipient.VerifyToken = values.Mother().IDValid()

	s.Message = fmt.Sprintf("Hi %s, please confirm your account through this link: %s/verify?token=%s&id=%s", s.Recipient.Username.Value(), appServerURL, s.Recipient.VerifyToken.Value(), s.Recipient.ID.Value())

	s.Logger = new(records.LoggerMock)

	s.SUT = &terminal.Confirmation{
		Logger:       s.Logger,
		AppServerURL: appServerURL,
	}
}

func TestIntegrationConfirmationSuite(t *testing.T) {
	suite.Run(t, new(ConfirmationTestSuite))
}
