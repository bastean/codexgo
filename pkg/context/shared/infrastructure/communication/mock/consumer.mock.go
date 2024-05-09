package communicationMock

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/message"
	"github.com/bastean/codexgo/pkg/context/shared/domain/queue"
	"github.com/stretchr/testify/mock"
)

type ConsumerMock struct {
	mock.Mock
}

func (consumer *ConsumerMock) SubscribedTo() []*queue.Queue {
	args := consumer.Called()
	return args.Get(0).([]*queue.Queue)
}

func (consumer *ConsumerMock) On(message *message.Message) error {
	// TODO?(goroutine): consumer.Called(message)
	return nil
}
