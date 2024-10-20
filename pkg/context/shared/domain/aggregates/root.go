package aggregates

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
)

type Root struct {
	Messages []*messages.Message
}

func (root *Root) Record(message *messages.Message) {
	root.Messages = append(root.Messages, message)
}

func (root *Root) Pull() []*messages.Message {
	recorded := root.Messages

	root.Messages = []*messages.Message{}

	return recorded
}

func NewRoot() *Root {
	return &Root{
		Messages: []*messages.Message{},
	}
}
