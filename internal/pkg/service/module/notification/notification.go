package notification

import (
	"github.com/bastean/codexgo/v4/pkg/context/notification/application/confirmation"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events/user"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/transfers"
)

func Init(transfer transfers.Transfer[*user.CreatedSucceededAttributes]) {
	Confirmation = &confirmation.Consumer{
		Confirmation: &confirmation.Case{
			Transfer: transfer,
		},
	}
}
