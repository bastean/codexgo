package errors

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mother"
)

type m struct {
	*mother.Mother
}

func (m *m) BubbleValid() *Bubble {
	return &Bubble{
		When:  m.TimeNow(),
		Where: m.LoremIpsumWord(),
		What:  m.LoremIpsumSentence(m.IntRange(1, 3)),
		Why: Meta{
			m.LoremIpsumWord(): m.LoremIpsumSentence(m.IntRange(1, 3)),
		},
		Who: m.Error(),
	}
}

func (m *m) BubbleInvalidWithoutWhere() {
	func() {
		_ = New[Default](&Bubble{What: m.LoremIpsumWord()})
	}()
}

func (m *m) BubbleInvalidWithoutWhat() {
	_ = New[Default](&Bubble{Where: m.LoremIpsumWord()})
}

func (m *m) BubbleUpValid() (error, error) {
	err := m.Error()
	return BubbleUp(err), err
}

func (m *m) PanicValidWithError(err error) {
	Panic(err)
}

var Mother = mother.New[m]
