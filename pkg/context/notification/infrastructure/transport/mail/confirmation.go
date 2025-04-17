package mail

import (
	"bytes"
	"context"
	"fmt"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/transports/smtp"
)

type Confirmation struct {
	*smtp.SMTP
	AppServerURL string
}

func (c *Confirmation) Submit(attributes *events.UserCreatedSucceededAttributes) error {
	var message bytes.Buffer

	headers := c.Headers(attributes.Email, "Account Confirmation")

	_, err := message.Write([]byte(headers))

	if err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			What: "Failure to write message headers",
			Why: errors.Meta{
				"Headers":   headers,
				"Verify ID": attributes.Verify,
				"User ID":   attributes.ID,
			},
			Who: err,
		})
	}

	link := fmt.Sprintf("%s/v4/account/verify?token=%s&id=%s", c.AppServerURL, attributes.Verify, attributes.ID)

	err = ConfirmationTemplate(attributes.Username, link).Render(context.Background(), &message)

	if err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			What: "Failure to render account confirmation mail",
			Why: errors.Meta{
				"Verify ID": attributes.Verify,
				"User ID":   attributes.ID,
			},
			Who: err,
		})
	}

	err = c.SendMail([]string{attributes.Email}, message.Bytes())

	if err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			What: "Failure to send account confirmation mail",
			Why: errors.Meta{
				"SMTP Server URL": c.ServerURL,
				"Verify ID":       attributes.Verify,
				"User ID":         attributes.ID,
			},
			Who: err,
		})
	}

	return nil
}
