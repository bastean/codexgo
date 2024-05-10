package sendMail

import (
	"github.com/bastean/codexgo/pkg/context/notify/domain/model"
	"github.com/bastean/codexgo/pkg/context/shared/domain/serror"
	"github.com/bastean/codexgo/pkg/context/shared/domain/stype"
)

type SendMail struct {
	model.Mail
}

func (sendMail *SendMail) Run(mail model.MailTemplate) (*stype.Empty, error) {
	err := sendMail.Mail.Send(mail)

	if err != nil {
		return nil, serror.BubbleUp(err, "Run")
	}

	return nil, nil
}
