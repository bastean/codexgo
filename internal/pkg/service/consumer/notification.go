package consumer

import (
	"github.com/bastean/codexgo/v4/internal/pkg/adapter/event"
	"github.com/bastean/codexgo/v4/internal/pkg/adapter/log"
	"github.com/bastean/codexgo/v4/internal/pkg/service/env"
	"github.com/bastean/codexgo/v4/internal/pkg/service/transport"
	"github.com/bastean/codexgo/v4/pkg/context/notification/application/confirmation"
	"github.com/bastean/codexgo/v4/pkg/context/notification/application/password"
	"github.com/bastean/codexgo/v4/pkg/context/notification/domain/role"
	"github.com/bastean/codexgo/v4/pkg/context/notification/infrastructure/transport/mail"
	"github.com/bastean/codexgo/v4/pkg/context/notification/infrastructure/transport/terminal"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
)

var (
	NotificationConfirmation *confirmation.Consumer
	NotificationPassword     *password.Consumer
)

func InitNotification() error {
	var (
		transferConfirmation role.Transfer[*events.UserCreatedSucceededAttributes]
		transferPassword     role.Transfer[*events.UserResetQueuedAttributes]
	)

	switch {
	case env.HasSMTP():
		transferConfirmation = &mail.Confirmation{
			SMTP:         transport.SMTP,
			AppServerURL: env.ServerGinURL,
		}

		transferPassword = &mail.Password{
			SMTP:         transport.SMTP,
			AppServerURL: env.ServerGinURL,
		}
	default:
		transferConfirmation = &terminal.Confirmation{
			Logger:       log.Log,
			AppServerURL: env.ServerGinURL,
		}

		transferPassword = &terminal.Password{
			Logger:       log.Log,
			AppServerURL: env.ServerGinURL,
		}
	}

	NotificationConfirmation = &confirmation.Consumer{
		Confirmation: &confirmation.Case{
			Transfer: transferConfirmation,
		},
	}

	NotificationPassword = &password.Consumer{
		Password: &password.Case{
			Transfer: transferPassword,
		},
	}

	err = events.AddEventMapper(event.Bus, events.Mapper{
		events.UserCreatedSucceededKey: {
			NotificationConfirmation,
		},
		events.UserResetQueuedKey: {
			NotificationPassword,
		},
	})

	if err != nil {
		return errors.BubbleUp(err, "InitNotification")
	}

	return nil
}
