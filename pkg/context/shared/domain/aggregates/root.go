package saggregate

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/smessage"
)

type AggregateRoot struct {
	Messages []*smessage.Message
}

func (root *AggregateRoot) RecordMessage(message *smessage.Message) {
	root.Messages = append(root.Messages, message)
}

func (root *AggregateRoot) PullMessages() []*smessage.Message {
	recordedMessages := root.Messages

	root.Messages = []*smessage.Message{}

	return recordedMessages
}

func NewAggregateRoot() *AggregateRoot {
	return &AggregateRoot{
		Messages: []*smessage.Message{},
	}
}
