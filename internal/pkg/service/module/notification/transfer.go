package notification

import (
	"github.com/bastean/codexgo/v4/pkg/context/notification/infrastructure/transport/mail"
	"github.com/bastean/codexgo/v4/pkg/context/notification/infrastructure/transport/terminal"
)

type (
	MailConfirmation     = mail.Confirmation
	TerminalConfirmation = terminal.Confirmation
)
