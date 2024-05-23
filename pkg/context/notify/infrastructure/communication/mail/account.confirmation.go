package mail

import (
	"bytes"
	"context"
	"fmt"
	"net/smtp"

	"github.com/bastean/codexgo/pkg/context/notify/application/send"
	"github.com/bastean/codexgo/pkg/context/shared/domain/serror"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/stransport"
)

type AccountConfirmation struct {
	*stransport.SMTP
}

func (client *AccountConfirmation) Submit(data any) error {
	account := data.(*send.CreatedSucceededEventAttributes)

	var message bytes.Buffer

	headers := fmt.Sprintf("From: %s\n"+"To: %s\n"+"Subject: Account Confirmation", client.Username, account.Email)

	message.Write([]byte(fmt.Sprintf("%s\n%s\n", headers, client.MIMEHeaders)))

	confirmationLink := fmt.Sprintf("%s/verify/%s", client.ServerUrl, account.Id)

	AccountConfirmationTemplate(account.Username, confirmationLink).Render(context.Background(), &message)

	err := smtp.SendMail(client.SmtpServerUrl, client.Auth, client.Username, []string{account.Email}, message.Bytes())

	if err != nil {
		return serror.NewInternal(&serror.Bubble{
			Where: "Submit",
			What:  "failure to send a account confirmation mail",
			Why: serror.Meta{
				"AccountId":     account.Id,
				"SmtpServerUrl": client.SmtpServerUrl,
			},
			Who: err,
		})
	}

	return nil
}
