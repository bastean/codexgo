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
	suite.sut = communication.NewNotifySmtpMail(suite.host, suite.port, suite.username, suite.password)
}

func (suite *SmtpMailTestSuite) TestSendAccountConfirmation() {
	mailTemplate := template.NewMail([]string{suite.username})

	accountConfirmationTemplate := template.NewAccountConfirmationMail(mailTemplate, suite.username, "test-send-account-confirmation-success")

	suite.NotPanics(func() { suite.sut.Send(accountConfirmationTemplate) })
}

func TestIntegrationSmtpMailSuite(t *testing.T) {
	suite.Run(t, new(SmtpMailTestSuite))
}
