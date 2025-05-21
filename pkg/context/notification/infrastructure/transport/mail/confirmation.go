package mail

import (
	"bytes"
	"context"
	"fmt"

	"github.com/bastean/codexgo/v4/pkg/context/notification/domain/aggregate/recipient"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/transports/smtp"
)

type Confirmation struct {
	*smtp.SMTP
	AppServerURL string
}

func (c *Confirmation) Submit(recipient *recipient.Recipient) error {
	var message bytes.Buffer

	headers := c.Headers(recipient.Email.Value(), "Account Confirmation")

	_, err := message.Write([]byte(headers))

	if err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			What: "Failure to write message headers",
			Why: errors.Meta{
				"Headers":   headers,
				"Verify ID": recipient.VerifyToken.Value(),
				"User ID":   recipient.ID.Value(),
			},
			Who: err,
		})
	}

	link := fmt.Sprintf("%s/v4/account/verify?token=%s&id=%s", c.AppServerURL, recipient.VerifyToken.Value(), recipient.ID.Value())

	err = ConfirmationTemplate(recipient.Username.Value(), link).Render(context.Background(), &message)

	if err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			What: "Failure to render account confirmation mail",
			Why: errors.Meta{
				"Verify ID": recipient.VerifyToken.Value(),
				"User ID":   recipient.ID.Value(),
			},
			Who: err,
		})
	}

	err = c.SendMail([]string{recipient.Email.Value()}, message.Bytes())

	if err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			What: "Failure to send account confirmation mail",
			Why: errors.Meta{
				"SMTP Server URL": c.ServerURL,
				"Verify ID":       recipient.VerifyToken.Value(),
				"User ID":         recipient.ID.Value(),
			},
			Who: err,
		})
	}

	return nil
}
