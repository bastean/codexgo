package user

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/pkg/context/shared/domain/transfers"
	"github.com/bastean/codexgo/pkg/context/user/application/created"
)

var (
	Created *created.Consumer
)

func InitCreated(transfer transfers.Transfer, queue *messages.Queue) {
	Created = &created.Consumer{
		Created: &created.Created{
			Transfer: transfer,
		},
		Queues: []*messages.Queue{queue},
	}
}
