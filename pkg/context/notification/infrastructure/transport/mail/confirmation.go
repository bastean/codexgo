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

func (client *Confirmation) Submit(data *user.CreatedSucceededAttributes) error {
	var message bytes.Buffer

	headers := client.Headers(data.Email, "Account Confirmation")

	_, err := message.Write([]byte(headers))

	if err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			Where: "Submit",
			What:  "Failure to write message headers",
			Why: errors.Meta{
				"Headers": headers,
				"User ID": data.ID,
			},
			Who: err,
		})
	}

	link := fmt.Sprintf("%s/v4/account/verify/%s", client.AppServerURL, data.ID)

	ConfirmationTemplate(data.Username, link).Render(context.Background(), &message)

	err = client.SendMail([]string{data.Email}, message.Bytes())

	if err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			Where: "Submit",
			What:  "Failure to send an account confirmation mail",
			Why: errors.Meta{
				"User ID":         data.ID,
				"SMTP Server URL": client.ServerURL,
			},
			Who: err,
		})
	}

	return nil
}
