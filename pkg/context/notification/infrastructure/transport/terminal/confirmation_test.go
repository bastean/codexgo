package terminal_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/notification/infrastructure/transport"
	"github.com/bastean/codexgo/v4/pkg/context/notification/infrastructure/transport/terminal"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events/user"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/records"
)

type ConfirmationTestSuite struct {
	transport.OfflineSuite[*user.CreatedSucceededAttributes]
}

func (s *ConfirmationTestSuite) SetupTest() {
	appServerURL := os.Getenv("CODEXGO_SERVER_GIN_URL")

	s.OfflineSuite.Attributes = new(user.CreatedSucceededAttributes)

	messages.RandomizeAttributes(s.OfflineSuite.Attributes)

	s.OfflineSuite.Message = fmt.Sprintf("Hi %s, please confirm your account through this link: %s/v4/account/verify/%s", s.OfflineSuite.Attributes.Username, appServerURL, s.OfflineSuite.Attributes.ID)

	s.OfflineSuite.Logger = new(records.LoggerMock)

	s.OfflineSuite.SUT = &terminal.Confirmation{
		Logger:       s.OfflineSuite.Logger,
		AppServerURL: appServerURL,
	}
}

func TestIntegrationConfirmationSuite(t *testing.T) {
	suite.Run(t, new(ConfirmationTestSuite))
}
