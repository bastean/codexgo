package send

import (
	"github.com/bastean/codexgo/pkg/context/notify/domain/model"
	"github.com/bastean/codexgo/pkg/context/shared/domain/serror"
	"github.com/bastean/codexgo/pkg/context/shared/domain/stype"
)

type Send struct {
	model.Transport
}

func (send *Send) Run(data any) (*stype.Empty, error) {
	err := send.Transport.Submit(data)

	if err != nil {
		return nil, serror.BubbleUp(err, "Run")
	}

	return nil, nil
}
