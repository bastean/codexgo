package terminal

import (
	"fmt"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/user/domain/event"
)

type Confirmation struct {
	models.Logger
	ServerURL string
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

	link := fmt.Sprintf("Hi %s, please confirm your account through this link: %s/verify/%s", user.Username, client.ServerURL, user.Id)

	client.Logger.Info(link)

	return nil
}
