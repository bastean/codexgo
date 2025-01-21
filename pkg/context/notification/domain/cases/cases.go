package cases

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
)

type (
	Confirmation interface {
		Run(*events.UserCreatedSucceededAttributes) error
	}
	Password interface {
		Run(*events.UserResetQueuedAttributes) error
	}
)
