package mail_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/transfers"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/transports/smtp"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/transport/mail"
	"github.com/stretchr/testify/suite"
)

type ConfirmationTestSuite struct {
	suite.Suite
	sut  transfers.Transfer
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
	message := user.RandomCreatedSucceeded()

	event := new(user.CreatedSucceeded)

	event.Attributes = new(user.CreatedSucceededAttributes)

	suite.NoError(json.Unmarshal(message.Attributes, event.Attributes))

	suite.NoError(suite.sut.Submit(event.Attributes))
}

func TestIntegrationConfirmationSuite(t *testing.T) {
	suite.Run(t, new(ConfirmationTestSuite))
}
