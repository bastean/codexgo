package mail_test

import (
	"os"
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/notification/domain/aggregate/recipient"
	"github.com/bastean/codexgo/v4/pkg/context/notification/infrastructure/transport"
	"github.com/bastean/codexgo/v4/pkg/context/notification/infrastructure/transport/mail"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/transports/smtp"
)

type PasswordTestSuite struct {
	transport.OnlineSuite
}

func (s *PasswordTestSuite) SetupSuite() {
	s.Recipient = recipient.Mother().RecipientValid()

	s.Recipient.ResetToken = values.Mother().TokenValid()

	smtp := smtp.Open(
		&smtp.Auth{
			Host:     os.Getenv("CODEXGO_SMTP_HOST"),
			Port:     os.Getenv("CODEXGO_SMTP_PORT"),
			Username: os.Getenv("CODEXGO_SMTP_USERNAME"),
			Password: os.Getenv("CODEXGO_SMTP_PASSWORD"),
		},
	)

	s.SUT = &mail.Password{
		SMTP:         smtp,
		AppServerURL: os.Getenv("CODEXGO_SERVER_GIN_URL"),
	}
}

func TestIntegrationPasswordSuite(t *testing.T) {
	suite.Run(t, new(PasswordTestSuite))
}
