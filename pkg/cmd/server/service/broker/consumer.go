package broker

import (
	"github.com/bastean/codexgo/pkg/cmd/server/service/notify"
	"github.com/bastean/codexgo/pkg/context/notify/application/send"
	"github.com/bastean/codexgo/pkg/context/shared/domain/queues"
)

var NotifySendAccountConfirmationQueueConsumer = &send.CreatedSucceededEventConsumer{
	UseCase: notify.SendAccountConfirmation,
	Queues:  []*queues.Queue{NotifySendAccountConfirmationQueue},
}
