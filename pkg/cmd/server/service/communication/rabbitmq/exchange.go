package rabbitmq

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/messages"
)

var Exchange = &messages.Router{
	Name: "codexgo",
}
