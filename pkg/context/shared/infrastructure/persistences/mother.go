package persistences

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mother"
)

type m struct {
	*mother.Mother
}

func (m *m) KeyValuesValid() []string {
	return m.Words(m.IntRange(1, 12))
}

var Mother = mother.New[m]()
