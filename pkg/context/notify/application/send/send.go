package send

import (
	"github.com/bastean/codexgo/pkg/context/notify/domain/model"
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/types"
)

type Send struct {
	model.Transport
}

func (send *Send) Run(data any) (types.Empty, error) {
	err := send.Transport.Submit(data)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return nil, nil
}
