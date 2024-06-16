package rabbitmq

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/domain/queues"
	"github.com/bastean/codexgo/pkg/context/user/application/created"
)

var UserSendConfirmationQueueConsumer = &created.Consumer{
	Queues: []*queues.Queue{UserSendConfirmationQueue},
}

func Consumers(created *created.Created) []models.Consumer {
	UserSendConfirmationQueueConsumer.UseCase = created

	return []models.Consumer{
		UserSendConfirmationQueueConsumer,
	}
}
