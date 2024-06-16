package created

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/domain/types"
	"github.com/bastean/codexgo/pkg/context/user/domain/event"
)

type Created struct {
	models.Transport
}

func (created *Created) Run(user *event.CreatedSucceeded) (types.Empty, error) {
	err := created.Transport.Submit(user.Attributes)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return nil, nil
}
