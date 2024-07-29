package rabbitmq

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
)

func Exchange(name string) *messages.Router {
	return &messages.Router{
		Name: name,
	}
}
