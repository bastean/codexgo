package communication_test

import (
	"os"
	"testing"

	"github.com/bastean/codexgo/pkg/context/notify/domain/model"
	"github.com/bastean/codexgo/pkg/context/notify/domain/template"
	"github.com/bastean/codexgo/pkg/context/notify/infrastructure/communication"
	"github.com/stretchr/testify/suite"
)

type SmtpMailTestSuite struct {
	suite.Suite
	sut                            model.Mail
	host, port, username, password string
}

func (suite *SmtpMailTestSuite) SetupTest() {
	suite.host = os.Getenv("SMTP_HOST")
	suite.port = os.Getenv("SMTP_PORT")
	suite.username = os.Getenv("SMTP_USERNAME")
	suite.password = os.Getenv("SMTP_PASSWORD")

	serverUrl := os.Getenv("URL")

	suite.sut = communication.NewNotifySmtpMail(suite.host, suite.port, suite.username, suite.password, serverUrl)
}

func (suite *SmtpMailTestSuite) TestSendAccountConfirmation() {
	accountConfirmationTemplate := &template.AccountConfirmationMail{
		Mail: &template.Mail{
			To: []string{suite.username},
		},
		Username:         suite.username,
		ConfirmationLink: "test-send-account-confirmation-success",
	}

	suite.NoError(suite.sut.Send(accountConfirmationTemplate))
}

func TestIntegrationSmtpMailSuite(t *testing.T) {
	suite.Run(t, new(SmtpMailTestSuite))
}
