package aggregates

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
)

type Root struct {
	Messages []*messages.Message
}

func (root *Root) RecordMessage(message *messages.Message) {
	root.Messages = append(root.Messages, message)
}

func (root *Root) PullMessages() []*messages.Message {
	recordedMessages := root.Messages

	root.Messages = []*messages.Message{}

	return recordedMessages
}

func NewRoot() *Root {
	return &Root{
		Messages: []*messages.Message{},
	}
}
