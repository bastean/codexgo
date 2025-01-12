package user

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
)

var (
	CreatedSucceededKey = events.UserCreatedSucceededKey
)

type (
	CreatedSucceededAttributes = events.UserCreatedSucceededAttributes
	CreatedSucceededMeta       = events.UserCreatedSucceededMeta
)
