package mail

import (
	"bytes"
	"context"
	"fmt"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/transports/smtp"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate/user"
)

type Confirmation struct {
	*smtp.SMTP
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

	headers := client.Headers(attributes.Email, "Account Confirmation")

	_, err := message.Write([]byte(headers))

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Submit",
			What:  "failure to write message headers",
			Why: errors.Meta{
				"Headers": headers,
				"User Id": attributes.Id,
			},
			Who: err,
		})
	}

	link := fmt.Sprintf("%s/verify/%s", client.ServerURL, attributes.Id)

	ConfirmationTemplate(attributes.Username, link).Render(context.Background(), &message)

	err = client.SendMail([]string{attributes.Email}, message.Bytes())

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
