package user

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/user/application/created"
)

var (
	Created *created.Consumer
)

func InitCreated(transport models.Transport, queue *messages.Queue) {
	Created = &created.Consumer{
		Created: &created.Created{
			Transport: transport,
		},
		Queues: []*messages.Queue{queue},
	}
}
