package communications

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/pkg/context/shared/domain/queues"
	"github.com/stretchr/testify/mock"
)

type ConsumerMock struct {
	mock.Mock
}

func (consumer *ConsumerMock) SubscribedTo() []*queues.Queue {
	args := consumer.Called()
	return args.Get(0).([]*queues.Queue)
}

func (consumer *ConsumerMock) On(message *messages.Message) error {
	// TODO?(goroutine): consumer.Called(message)
	return nil
}
