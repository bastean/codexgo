package terminal

import (
	"fmt"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/loggers"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
)

type Confirmation struct {
	loggers.Logger
	ServerURL string
}

func (client *Confirmation) Submit(data any) error {
	attributes, ok := data.(*user.CreatedSucceededAttributes)

	if !ok {
		return errors.NewInternal(&errors.Bubble{
			Where: "Submit",
			What:  "Failure in type assertion",
			Why: errors.Meta{
				"Expected": new(user.CreatedSucceededAttributes),
				"Actual":   data,
			},
		})
	}

	link := fmt.Sprintf("Hi %s, please confirm your account through this link: %s/v4/account/verify/%s", attributes.Username, client.ServerURL, attributes.Id)

	client.Logger.Info(link)

	return nil
}
