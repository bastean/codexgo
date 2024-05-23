package notify

import (
	"os"

	"github.com/bastean/codexgo/pkg/cmd/server/service/logger"
	"github.com/bastean/codexgo/pkg/context/notify/application/send"
	"github.com/bastean/codexgo/pkg/context/notify/infrastructure/communication/mail"
	"github.com/bastean/codexgo/pkg/context/notify/infrastructure/communication/terminal"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/stransport"
)

var host = os.Getenv("SMTP_HOST")
var port = os.Getenv("SMTP_PORT")
var username = os.Getenv("SMTP_USERNAME")
var password = os.Getenv("SMTP_PASSWORD")
var serverURL = os.Getenv("URL")

var smtp = stransport.NewSMTP(host, port, username, password, serverURL)

var viaMail = &mail.AccountConfirmation{
	SMTP: smtp,
}

var viaTerminal = &terminal.AccountConfirmation{
	ServerURL: serverURL,
	Logger:    logger.Logger,
}

var SendAccountConfirmation = new(send.Send)

func Init() error {
	logger.Info("starting module: notify")

	if host != "" {
		SendAccountConfirmation.Transport = viaMail
	} else {
		SendAccountConfirmation.Transport = viaTerminal
	}

	return nil
}
