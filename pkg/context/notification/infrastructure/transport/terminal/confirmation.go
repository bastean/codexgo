package terminal

import (
	"fmt"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events/user"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/loggers"
)

type Confirmation struct {
	loggers.Logger
	AppServerURL string
}

func (client *Confirmation) Submit(data *user.CreatedSucceededAttributes) error {
	link := fmt.Sprintf("Hi %s, please confirm your account through this link: %s/v4/account/verify/%s", data.Username, client.AppServerURL, data.ID)

	client.Logger.Info(link)

	return nil
}
