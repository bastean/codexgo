package created

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate/user"
)

type Created struct {
	models.Transport
}

func (created *Created) Run(event *user.CreatedSucceeded) error {
	err := created.Transport.Submit(event.Attributes)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}
