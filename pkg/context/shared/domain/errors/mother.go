package errors

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mother"
)

type m struct {
	*mother.Mother
}

func (m *m) BubbleValid() *Bubble {
	return &Bubble{
		Where: m.LoremIpsumWord(),
		What:  m.LoremIpsumSentence(m.IntRange(1, 3)),
		Why: Meta{
			m.LoremIpsumWord(): m.LoremIpsumSentence(m.IntRange(1, 3)),
		},
		Who: m.Error(),
	}
}

func (m *m) BubbleValidWithoutWhere() error {
	return func() error {
		return New[Default](&Bubble{What: m.LoremIpsumWord()})
	}()
}

func (m *m) BubbleInvalidWithoutWhat() {
	_ = New[Default](&Bubble{Where: m.LoremIpsumWord()})
}

func (m *m) BubbleInvalidWhy() {
	_ = New[Default](&Bubble{
		What: m.LoremIpsumWord(),
		Why: Meta{
			m.LoremIpsumWord(): func() {},
		},
	}).Error()
}

func (m *m) DefaultValid() *Default {
	return New[Default](m.BubbleValid())
}

func (m *m) InternalValid() *Internal {
	return New[Internal](m.BubbleValid())
}

func (m *m) FailureValid() *Failure {
	return New[Failure](m.BubbleValid())
}

func (m *m) InvalidValueValid() *InvalidValue {
	return New[InvalidValue](m.BubbleValid())
}

func (m *m) AlreadyExistValid() *AlreadyExist {
	return New[AlreadyExist](m.BubbleValid())
}

func (m *m) NotExistValid() *NotExist {
	return New[NotExist](m.BubbleValid())
}

func (m *m) BubbleUpValid() (bubble error, value error) {
	err := m.Error()
	return BubbleUp(err), err
}

func (m *m) PanicValidWithError(err error) {
	Panic(err)
}

var Mother = mother.New[m]
