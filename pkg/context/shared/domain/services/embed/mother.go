package embed

import (
	"fmt"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mother"
)

type m struct {
	*mother.Mother
}

func (m *m) EmbedValid() (message, value string) {
	value = m.Word()
	return fmt.Sprintf("%s [%s]", m.WordsJoin(m.Words(m.IntRange(1, 3)), " "), value), value
}

func (m *m) EmbedInvalid() string {
	return m.RandomString(m.Words(m.IntRange(0, 3)))
}

var Mother = mother.New[m]
