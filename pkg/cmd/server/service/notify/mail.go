package notify

import (
	"github.com/bastean/codexgo/pkg/context/notify/domain/model"
	"github.com/bastean/codexgo/pkg/context/notify/infrastructure/communication/mail"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/transports"
)

func NewMailAccountConfirmation(smtp *transports.SMTP) model.Transport {
	return &mail.AccountConfirmation{
		SMTP: smtp,
	}
}
