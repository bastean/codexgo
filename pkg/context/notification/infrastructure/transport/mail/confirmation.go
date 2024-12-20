package mail

import (
	"bytes"
	"context"
	"fmt"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events/user"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/transports/smtp"
)

type Confirmation struct {
	*smtp.SMTP
	AppServerURL string
}

func (c *Confirmation) Submit(attributes *user.CreatedSucceededAttributes) error {
	var message bytes.Buffer

	headers := c.Headers(attributes.Email, "Account Confirmation")

	_, err := message.Write([]byte(headers))

	if err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			Where: "Submit",
			What:  "Failure to write message headers",
			Why: errors.Meta{
				"Headers": headers,
				"User ID": attributes.ID,
			},
			Who: err,
		})
	}

	link := fmt.Sprintf("%s/v4/account/verify/%s", c.AppServerURL, attributes.ID)

	ConfirmationTemplate(attributes.Username, link).Render(context.Background(), &message)

	err = c.SendMail([]string{attributes.Email}, message.Bytes())

	if err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			Where: "Submit",
			What:  "Failure to send an account confirmation mail",
			Why: errors.Meta{
				"User ID":         attributes.ID,
				"SMTP Server URL": c.ServerURL,
			},
			Who: err,
		})
	}

	return nil
}
