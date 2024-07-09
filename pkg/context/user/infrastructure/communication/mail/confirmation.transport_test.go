package mail_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/transports"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/communication/mail"
	"github.com/stretchr/testify/suite"
)

type MailConfirmationTransportTestSuite struct {
	suite.Suite
	sut  models.Transport
	smtp *transports.SMTP
}

func (suite *MailConfirmationTransportTestSuite) SetupTest() {
	suite.smtp = transports.NewSMTP(
		os.Getenv("CODEXGO_SMTP_HOST"),
		os.Getenv("CODEXGO_SMTP_PORT"),
		os.Getenv("CODEXGO_SMTP_USERNAME"),
		os.Getenv("CODEXGO_SMTP_PASSWORD"),
		os.Getenv("CODEXGO_SERVER_URL"),
	)

	suite.sut = &mail.Confirmation{
		SMTP: suite.smtp,
	}
}

func (suite *MailConfirmationTransportTestSuite) TestSubmit() {
	message := user.RandomCreatedSucceeded()

	event := new(user.CreatedSucceeded)

	event.Attributes = new(user.CreatedSucceededAttributes)

	suite.NoError(json.Unmarshal(message.Attributes, event.Attributes))

	suite.NoError(suite.sut.Submit(event.Attributes))
}

func TestIntegrationMailConfirmationTransportSuite(t *testing.T) {
	suite.Run(t, new(MailConfirmationTransportTestSuite))
}
