package notify

import (
	"os"

	"github.com/bastean/codexgo/pkg/context/notify/application/sendMail"
	"github.com/bastean/codexgo/pkg/context/notify/infrastructure/communication"
)

var host = os.Getenv("SMTP_HOST")
var port = os.Getenv("SMTP_PORT")
var username = os.Getenv("SMTP_USERNAME")
var password = os.Getenv("SMTP_PASSWORD")
var serverUrl = os.Getenv("URL")

var SMTP = communication.NewSMTP(host, port, username, password, serverUrl)

var SendMail = &sendMail.SendMail{
	Mail: SMTP,
}
