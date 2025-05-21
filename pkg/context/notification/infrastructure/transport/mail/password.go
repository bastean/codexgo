package mail

import (
	"bytes"
	"context"
	"fmt"

	"github.com/bastean/codexgo/v4/pkg/context/notification/domain/aggregate/recipient"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/transports/smtp"
)

type Password struct {
	*smtp.SMTP
	AppServerURL string
}

func (c *Password) Submit(recipient *recipient.Recipient) error {
	var message bytes.Buffer

	headers := c.Headers(recipient.Email.Value(), "Password Reset")

	_, err := message.Write([]byte(headers))

	if err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			What: "Failure to write message headers",
			Why: errors.Meta{
				"Headers":  headers,
				"Reset ID": recipient.ResetToken.Value(),
				"User ID":  recipient.ID.Value(),
			},
			Who: err,
		})
	}

	link := fmt.Sprintf("%s/reset?token=%s&id=%s", c.AppServerURL, recipient.ResetToken.Value(), recipient.ID.Value())

	err = PasswordTemplate(recipient.Username.Value(), link).Render(context.Background(), &message)

	if err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			What: "Failure to render password reset mail",
			Why: errors.Meta{
				"Reset ID": recipient.ResetToken.Value(),
				"User ID":  recipient.ID.Value(),
			},
			Who: err,
		})
	}

	err = c.SendMail([]string{recipient.Email.Value()}, message.Bytes())

	if err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			What: "Failure to send password reset mail",
			Why: errors.Meta{
				"SMTP Server URL": c.ServerURL,
				"Reset ID":        recipient.ResetToken.Value(),
				"User ID":         recipient.ID.Value(),
			},
			Who: err,
		})
	}

	return nil
}
