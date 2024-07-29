package created

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/transfers"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
)

type Created struct {
	transfers.Transfer
}

func (created *Created) Run(event *user.CreatedSucceeded) error {
	err := created.Transfer.Submit(event.Attributes)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}
