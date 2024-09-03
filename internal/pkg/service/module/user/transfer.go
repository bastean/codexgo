package user

import (
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/transport/mail"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/transport/terminal"
)

type (
	MailConfirmation     = mail.Confirmation
	TerminalConfirmation = terminal.Confirmation
)
