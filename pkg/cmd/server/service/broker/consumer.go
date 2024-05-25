package broker

import (
	"github.com/bastean/codexgo/pkg/cmd/server/service/notify"
	"github.com/bastean/codexgo/pkg/context/notify/application/send"
	"github.com/bastean/codexgo/pkg/context/shared/domain/squeue"
)

var NotifySendAccountConfirmationQueueConsumer = &send.CreatedSucceededEventConsumer{
	UseCase: notify.SendAccountConfirmation,
	Queues:  []*squeue.Queue{NotifySendAccountConfirmationQueue},
}
