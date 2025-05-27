package aggregates

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mother"
)

type m struct {
	*mother.Mother
}

func (m *m) RootValid() *Root {
	return &Root{
		Events: make([]*messages.Message, 0),
	}
}

var Mother = mother.New[m]
