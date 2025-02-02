package consumer

import (
	"github.com/bastean/codexgo/v4/internal/pkg/adapter/event"
	"github.com/bastean/codexgo/v4/internal/pkg/adapter/log"
	"github.com/bastean/codexgo/v4/internal/pkg/service/env"
	"github.com/bastean/codexgo/v4/internal/pkg/service/transport"
	"github.com/bastean/codexgo/v4/pkg/context/notification/application/confirmation"
	"github.com/bastean/codexgo/v4/pkg/context/notification/domain/role"
	"github.com/bastean/codexgo/v4/pkg/context/notification/infrastructure/transport/mail"
	"github.com/bastean/codexgo/v4/pkg/context/notification/infrastructure/transport/terminal"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
)

var (
	NotificationConfirmation *confirmation.Consumer
)

func InitNotification() error {
	var (
		transfer role.Transfer[*events.UserCreatedSucceededAttributes]
	)

	switch {
	case env.HasSMTP():
		transfer = &mail.Confirmation{
			SMTP:         transport.SMTP,
			AppServerURL: env.ServerGinURL,
		}
	default:
		transfer = &terminal.Confirmation{
			Logger:       log.Log,
			AppServerURL: env.ServerGinURL,
		}
	}

	NotificationConfirmation = &confirmation.Consumer{
		Confirmation: &confirmation.Case{
			Transfer: transfer,
		},
	}

	err = events.AddEventMapper(event.Bus, events.Mapper{
		events.UserCreatedSucceededKey: {
			NotificationConfirmation,
		},
	})

	if err != nil {
		return errors.BubbleUp(err, "InitNotification")
	}

	return nil
}
