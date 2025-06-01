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

type PasswordTestSuite struct {
	transport.OfflineSuite
}

func (s *PasswordTestSuite) SetupSuite() {
	appServerURL := os.Getenv("CODEXGO_SERVER_GIN_URL")

	s.Recipient = recipient.Mother().RecipientValid()

	s.Recipient.ResetToken = values.Mother().IDValid()

	s.Message = fmt.Sprintf("Hi %s, please reset your password through this link: %s/reset?token=%s&id=%s", s.Recipient.Username.Value(), appServerURL, s.Recipient.ResetToken.Value(), s.Recipient.ID.Value())

	s.Logger = new(records.LoggerMock)

	s.SUT = &terminal.Password{
		Logger:       s.Logger,
		AppServerURL: appServerURL,
	}
}

func TestIntegrationPasswordSuite(t *testing.T) {
	suite.Run(t, new(PasswordTestSuite))
}
