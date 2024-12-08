package notification

import (
	"github.com/bastean/codexgo/v4/internal/pkg/service/module/user"
	"github.com/bastean/codexgo/v4/pkg/context/notification/application/confirmation"
	"github.com/bastean/codexgo/v4/pkg/context/notification/domain/transfer"
)

func Init(transfer transfer.Transfer[*user.CreatedSucceededAttributes]) {
	Confirmation = &confirmation.Consumer{
		Confirmation: &confirmation.Case{
			Transfer: transfer,
		},
	}
}
