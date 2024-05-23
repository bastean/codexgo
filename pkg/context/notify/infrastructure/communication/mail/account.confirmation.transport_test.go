package mail_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/bastean/codexgo/pkg/context/notify/application/send"
	"github.com/bastean/codexgo/pkg/context/notify/domain/model"
	"github.com/bastean/codexgo/pkg/context/notify/infrastructure/communication/mail"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/stransport"
	"github.com/stretchr/testify/suite"
)

type NotifyMailAccountConfirmationTransportTestSuite struct {
	suite.Suite
	sut                            model.Transport
	smtp                           *stransport.SMTP
	host, port, username, password string
}

func (suite *NotifyMailAccountConfirmationTransportTestSuite) SetupTest() {
	suite.host = os.Getenv("SMTP_HOST")
	suite.port = os.Getenv("SMTP_PORT")
	suite.username = os.Getenv("SMTP_USERNAME")
	suite.password = os.Getenv("SMTP_PASSWORD")

	serverUrl := os.Getenv("URL")

	suite.smtp = stransport.NewSMTP(
		suite.host,
		suite.port,
		suite.username,
		suite.password,
		serverUrl,
	)

	suite.sut = &mail.AccountConfirmation{
		SMTP: suite.smtp,
	}
}

func (suite *NotifyMailAccountConfirmationTransportTestSuite) TestSubmit() {
	message := send.RandomEvent()

	attributes := new(send.CreatedSucceededEventAttributes)

	json.Unmarshal(message.Attributes, attributes)

	suite.NoError(suite.sut.Submit(attributes))
}

func TestIntegrationNotifyMailAccountConfirmationTransportSuite(t *testing.T) {
	suite.Run(t, new(NotifyMailAccountConfirmationTransportTestSuite))
}
