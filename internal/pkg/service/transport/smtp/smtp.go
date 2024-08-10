package smtp

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/transports/smtp"
)

type (
	SMTP = smtp.SMTP
	Auth = smtp.Auth
)

var (
	Open = smtp.Open
)
