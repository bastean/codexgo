package cases

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events/user"
)

type (
	Confirmation interface {
		Run(*user.CreatedSucceededAttributes) error
	}
)
