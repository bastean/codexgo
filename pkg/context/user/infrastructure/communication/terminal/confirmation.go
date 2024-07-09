package terminal

import (
	"fmt"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate/user"
)

type Confirmation struct {
	models.Logger
	ServerURL string
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

	link := fmt.Sprintf("Hi %s, please confirm your account through this link: %s/verify/%s", attributes.Username, client.ServerURL, attributes.Id)

	client.Logger.Info(link)

	return nil
}
