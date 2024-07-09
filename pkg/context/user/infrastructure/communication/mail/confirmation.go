package mail

import (
	"bytes"
	"context"
	"fmt"
	"net/smtp"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/transports"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate/user"
)

type Confirmation struct {
	*transports.SMTP
}

func (client *Confirmation) Submit(data any) error {
	attributes, ok := data.(*user.CreatedSucceededAttributes)

	if !ok {
		return errors.NewInternal(&errors.Bubble{
			Where: "Submit",
			What:  "failure in type assertion",
			Why: errors.Meta{
				"Expected": new(user.CreatedSucceededAttributes),
				"Actual":   data,
			},
		})
	}

	var message bytes.Buffer

	headers := fmt.Sprintf("From: %s\n"+"To: %s\n"+"Subject: Account Confirmation", client.Username, attributes.Email)

	_, _ = message.Write([]byte(fmt.Sprintf("%s\n%s\n", headers, client.MIMEHeaders)))

	link := fmt.Sprintf("%s/verify/%s", client.ServerURL, attributes.Id)

	ConfirmationTemplate(attributes.Username, link).Render(context.Background(), &message)

	err := smtp.SendMail(client.SMTPServerURL, client.Auth, client.Username, []string{attributes.Email}, message.Bytes())

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Submit",
			What:  "failure to send an account confirmation mail",
			Why: errors.Meta{
				"User Id":         attributes.Id,
				"SMTP Server URL": client.SMTPServerURL,
			},
			Who: err,
		})
	}

	return nil
}
