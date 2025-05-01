package errors

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mother"
)

type m struct {
	*mother.Mother
}

func (m *m) BubbleInvalidWithoutWhere() *Bubble {
	return &Bubble{
		What: m.LoremIpsumWord(),
	}
}

func (m *m) BubbleInvalidWithoutWhat() *Bubble {
	return &Bubble{
		Where: m.LoremIpsumWord(),
	}
}

func (m *m) BubbleUpValid() (error, error) {
	err := m.Error()
	return BubbleUp(err), err
}

func (m *m) PanicValidWithError(err error) {
	Panic(err)
}

var Mother = mother.New[m]()
