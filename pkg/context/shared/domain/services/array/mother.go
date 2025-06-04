package array

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mother"
)

type m struct {
	*mother.Mother
}

func (m *m) SliceValid() ([]string, int) {
	values := make([]string, m.IntRange(3, 12))

	m.Slice(&values)

	return values, m.IntRange(0, len(values)-1)
}

func (m *m) SliceInvalid() ([]string, int) {
	return make([]string, 0), m.Int()
}

var Mother = mother.New[m]
