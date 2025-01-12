package notification

import (
	"github.com/bastean/codexgo/v4/pkg/context/notification/application/confirmation"
	"github.com/bastean/codexgo/v4/pkg/context/notification/domain/transfer"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
)

func Init(transfer transfer.Transfer[*events.UserCreatedSucceededAttributes]) {
	Confirmation = &confirmation.Consumer{
		Confirmation: &confirmation.Case{
			Transfer: transfer,
		},
	}
}
