package terminal_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/notification/infrastructure/transport"
	"github.com/bastean/codexgo/v4/pkg/context/notification/infrastructure/transport/terminal"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/records"
)

type PasswordTestSuite struct {
	transport.OfflineSuite[*events.UserResetQueuedAttributes]
}

func (s *PasswordTestSuite) SetupTest() {
	appServerURL := os.Getenv("CODEXGO_SERVER_GIN_URL")

	s.OfflineSuite.Attributes = new(events.UserResetQueuedAttributes)

	messages.RandomizeAttributes(s.OfflineSuite.Attributes)

	s.OfflineSuite.Message = fmt.Sprintf("Hi %s, please reset your password through this link: %s/reset?token=%s&id=%s", s.OfflineSuite.Attributes.Username, appServerURL, s.OfflineSuite.Attributes.Reset, s.OfflineSuite.Attributes.ID)

	s.OfflineSuite.Logger = new(records.LoggerMock)

	s.OfflineSuite.SUT = &terminal.Password{
		Logger:       s.OfflineSuite.Logger,
		AppServerURL: appServerURL,
	}
}

func TestIntegrationPasswordSuite(t *testing.T) {
	suite.Run(t, new(PasswordTestSuite))
}
