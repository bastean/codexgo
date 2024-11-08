package confirmation

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events/user"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/transfers"
)

type Confirmation struct {
	transfers.Transfer[*user.CreatedSucceededAttributes]
}

func (use *Confirmation) Run(event *user.CreatedSucceededAttributes) error {
	err := use.Transfer.Submit(event)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}
