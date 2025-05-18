package mail_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/notification/infrastructure/transport"
	"github.com/bastean/codexgo/v4/pkg/context/notification/infrastructure/transport/mail"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/transports/smtp"
)

type ConfirmationTestSuite struct {
	transport.OnlineSuite[*events.UserCreatedSucceededAttributes]
}

func (s *ConfirmationTestSuite) SetupSuite() {
	smtp := smtp.Open(
		&smtp.Auth{
			Host:     os.Getenv("CODEXGO_SMTP_HOST"),
			Port:     os.Getenv("CODEXGO_SMTP_PORT"),
			Username: os.Getenv("CODEXGO_SMTP_USERNAME"),
			Password: os.Getenv("CODEXGO_SMTP_PASSWORD"),
		},
	)

	s.OnlineSuite.Attributes = new(events.UserCreatedSucceededAttributes)

	transport.Mother().StructRandomize(s.OnlineSuite.Attributes)

	s.OnlineSuite.Attributes.Email = transport.Mother().Email()

	s.OnlineSuite.SUT = &mail.Confirmation{
		SMTP:         smtp,
		AppServerURL: os.Getenv("CODEXGO_SERVER_GIN_URL"),
	}
}

func TestIntegrationConfirmationSuite(t *testing.T) {
	suite.Run(t, new(ConfirmationTestSuite))
}
