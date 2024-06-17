package mail

import (
	"bytes"
	"context"
	"fmt"
	"net/smtp"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/transports"
	"github.com/bastean/codexgo/pkg/context/user/domain/event"
)

type Confirmation struct {
	*transports.SMTP
}

func (client *Confirmation) Submit(data any) error {
	user, ok := data.(*event.CreatedSucceededAttributes)

	if !ok {
		return errors.NewInternal(&errors.Bubble{
			Where: "Submit",
			What:  "failure in type assertion",
			Why: errors.Meta{
				"Expected": new(event.CreatedSucceededAttributes),
				"Actual":   data,
			},
		})
	}

	var message bytes.Buffer

	headers := fmt.Sprintf("From: %v\n"+"To: %v\n"+"Subject: Account Confirmation", client.Username, user.Email)

	_, _ = message.Write([]byte(fmt.Sprintf("%v\n%v\n", headers, client.MIMEHeaders)))

	link := fmt.Sprintf("%v/verify/%v", client.ServerURL, user.Id)

	ConfirmationTemplate(user.Username, link).Render(context.Background(), &message)

	err := smtp.SendMail(client.SMTPServerURL, client.Auth, client.Username, []string{user.Email}, message.Bytes())

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Submit",
			What:  "failure to send an account confirmation mail",
			Why: errors.Meta{
				"User Id":         user.Id,
				"SMTP Server URL": client.SMTPServerURL,
			},
			Who: err,
		})
	}

	return nil
}
