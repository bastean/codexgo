package notify

import (
	"github.com/bastean/codexgo/pkg/context/notify/application/sendMail"
	"github.com/bastean/codexgo/pkg/context/notify/infrastructure/communication"
)

var notifySmtpMail = communication.NewNotifySmtpMail()

var NotifySendMail = sendMail.NewSendMail(notifySmtpMail)
