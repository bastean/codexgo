package aggregate

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/message"
)

type AggregateRoot struct {
	Messages []*message.Message
}

func (root *AggregateRoot) RecordMessage(message *message.Message) {
	root.Messages = append(root.Messages, message)
}

func (root *AggregateRoot) PullMessages() []*message.Message {
	recordedMessages := root.Messages

	root.Messages = []*message.Message{}

	return recordedMessages
}

func NewAggregateRoot() *AggregateRoot {
	return &AggregateRoot{Messages: []*message.Message{}}
}
