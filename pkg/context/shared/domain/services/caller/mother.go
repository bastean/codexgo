package caller

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mother"
)

type m struct {
	*mother.Mother
}

func (m *m) ParseValidValues() (pkg, receiver, method string) {
	return m.LoremIpsumWord(), m.LoremIpsumWord(), m.LoremIpsumWord()
}

var Mother = mother.New[m]
