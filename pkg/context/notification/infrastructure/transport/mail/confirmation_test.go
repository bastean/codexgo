package mail_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/notification/infrastructure/transport"
	"github.com/bastean/codexgo/v4/pkg/context/notification/infrastructure/transport/mail"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events/user"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/transports/smtp"
)

type ConfirmationTestSuite struct {
	transport.OnlineSuite[*user.CreatedSucceededAttributes]
}

func (s *ConfirmationTestSuite) SetupTest() {
	smtp := smtp.Open(
		&smtp.Auth{
			Host:     os.Getenv("CODEXGO_SMTP_HOST"),
			Port:     os.Getenv("CODEXGO_SMTP_PORT"),
			Username: os.Getenv("CODEXGO_SMTP_USERNAME"),
			Password: os.Getenv("CODEXGO_SMTP_PASSWORD"),
		},
	)

	s.OnlineSuite.Attributes = new(user.CreatedSucceededAttributes)

	messages.RandomizeAttributes(s.Attributes)

	s.OnlineSuite.Attributes.Email = services.Create.Email()

	s.OnlineSuite.SUT = &mail.Confirmation{
		SMTP:         smtp,
		AppServerURL: os.Getenv("CODEXGO_SERVER_GIN_URL"),
	}
}

func TestIntegrationConfirmationSuite(t *testing.T) {
	suite.Run(t, new(ConfirmationTestSuite))
}
