package communications

import (
	"github.com/stretchr/testify/mock"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
)

type ConsumerMock struct {
	mock.Mock
}

func (consumer *ConsumerMock) SubscribedTo() []*messages.Queue {
	args := consumer.Called()
	return args.Get(0).([]*messages.Queue)
}

func (consumer *ConsumerMock) On(message *messages.Message) error {
	// TODO?(goroutine): consumer.Called(message)
	return nil
}
