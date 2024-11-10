package mail_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/notification/infrastructure/transport/mail"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events/user"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/transfers"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/transports/smtp"
)

type ConfirmationTestSuite struct {
	suite.Suite
	sut  transfers.Transfer[*user.CreatedSucceededAttributes]
	smtp *smtp.SMTP
}

func (suite *ConfirmationTestSuite) SetupTest() {
	suite.smtp = smtp.Open(
		&smtp.Auth{
			Host:     os.Getenv("CODEXGO_SMTP_HOST"),
			Port:     os.Getenv("CODEXGO_SMTP_PORT"),
			Username: os.Getenv("CODEXGO_SMTP_USERNAME"),
			Password: os.Getenv("CODEXGO_SMTP_PASSWORD"),
		},
	)

	suite.sut = &mail.Confirmation{
		SMTP:         suite.smtp,
		AppServerURL: os.Getenv("CODEXGO_SERVER_GIN_URL"),
	}
}

func (suite *ConfirmationTestSuite) TestSubmit() {
	attributes := new(user.CreatedSucceededAttributes)

	messages.RandomAttributes(attributes)

	attributes.Email = services.Create.Email()

	suite.NoError(suite.sut.Submit(attributes))
}

func TestIntegrationConfirmationSuite(t *testing.T) {
	suite.Run(t, new(ConfirmationTestSuite))
}
