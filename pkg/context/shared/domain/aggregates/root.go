package aggregates

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/messages"
)

type AggregateRoot struct {
	Messages []*messages.Message
}

func (root *AggregateRoot) RecordMessage(message *messages.Message) {
	root.Messages = append(root.Messages, message)
}

func (root *AggregateRoot) PullMessages() []*messages.Message {
	recordedMessages := root.Messages

	root.Messages = []*messages.Message{}

	return recordedMessages
}

func NewAggregateRoot() *AggregateRoot {
	return &AggregateRoot{
		Messages: []*messages.Message{},
	}
}
