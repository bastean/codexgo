package ascii

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mother"
)

type m struct {
	*mother.Mother
}

func (m *m) DrawingValid() (drawing []string, maxWidth int) {
	maxWidth = m.IntRange(3, 12)

	drawing = []string{
		m.Join(m.Letters(m.IntN(maxWidth)), ""),
		m.Join(m.Letters(maxWidth), ""),
		m.Join(m.Letters(m.IntN(maxWidth)), ""),
	}

	return drawing, maxWidth
}

var Mother = mother.New[m]
