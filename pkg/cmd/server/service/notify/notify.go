package notify

import (
	"os"

	"github.com/bastean/codexgo/pkg/context/notify/application/send"
	"github.com/bastean/codexgo/pkg/context/notify/infrastructure/communication/mail"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/stransport"
)

var host = os.Getenv("SMTP_HOST")
var port = os.Getenv("SMTP_PORT")
var username = os.Getenv("SMTP_USERNAME")
var password = os.Getenv("SMTP_PASSWORD")
var serverUrl = os.Getenv("URL")

var smtp = stransport.NewSMTP(host, port, username, password, serverUrl)

var accountConfirmationMail = &mail.AccountConfirmation{
	SMTP: smtp,
}

var SendAccountConfirmationMail = &send.Send{
	Transport: accountConfirmationMail,
}
