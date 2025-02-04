package mail

import (
	"bytes"
	"context"
	"fmt"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/transports/smtp"
)

type Password struct {
	*smtp.SMTP
	AppServerURL string
}

func (c *Password) Submit(attributes *events.UserResetQueuedAttributes) error {
	var message bytes.Buffer

	headers := c.Headers(attributes.Email, "Password Reset")

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

	link := fmt.Sprintf("%s/reset?token=%s&id=%s", c.AppServerURL, attributes.Reset, attributes.ID)

	PasswordTemplate(attributes.Username, link).Render(context.Background(), &message)

	err = c.SendMail([]string{attributes.Email}, message.Bytes())

	if err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			Where: "Submit",
			What:  "Failure to send password reset mail",
			Why: errors.Meta{
				"Reset ID":        attributes.Reset,
				"User ID":         attributes.ID,
				"SMTP Server URL": c.ServerURL,
			},
			Who: err,
		})
	}

	return nil
}
