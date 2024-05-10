package scommunication

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/smessage"
	"github.com/bastean/codexgo/pkg/context/shared/domain/squeue"
	"github.com/stretchr/testify/mock"
)

type ConsumerMock struct {
	mock.Mock
}

func (consumer *ConsumerMock) SubscribedTo() []*squeue.Queue {
	args := consumer.Called()
	return args.Get(0).([]*squeue.Queue)
}

func (consumer *ConsumerMock) On(message *smessage.Message) error {
	// TODO?(goroutine): consumer.Called(message)
	return nil
}
