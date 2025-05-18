package aggregates

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mother"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/time"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

type m struct {
	*mother.Mother
}

func (m *m) RootValid() *Root {
	return &Root{
		Events: make([]*messages.Message, 0),
	}
}

func (m *m) TimeValid() *Time {
	value, err := values.New[*Time](time.Now().Format())

	if err != nil {
		errors.Panic(err)
	}

	return value
}

func (m *m) TimeInvalid() (string, error) {
	var value string

	_, err := values.New[*Time](value)

	return value, err
}

var Mother = mother.New[m]()
