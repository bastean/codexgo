package rabbitmq

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/messages"
)

func Exchange(name string) *messages.Router {
	return &messages.Router{
		Name: name,
	}
}
