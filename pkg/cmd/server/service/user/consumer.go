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
	usecase := &created.Created{
		Transport: transport,
	}

	Created = &created.Consumer{
		UseCase: usecase,
		Queues:  []*messages.Queue{queue},
	}
}
