package sendMail

import (
	"github.com/bastean/codexgo/pkg/context/notify/domain/model"
	"github.com/bastean/codexgo/pkg/context/shared/domain/errs"
	"github.com/bastean/codexgo/pkg/context/shared/domain/types"
)

type SendMail struct {
	model.Mail
}

func (sendMail *SendMail) Run(mail model.MailTemplate) (*types.Empty, error) {
	err := sendMail.Mail.Send(mail)

	if err != nil {
		return nil, errs.BubbleUp(err, "Run")
	}

	return nil, nil
}
