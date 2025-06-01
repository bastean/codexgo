package password

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mother"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

type m struct {
	*mother.Mother
}

func (m *m) EventAttributesValid() *EventAttributes {
	return &EventAttributes{
		ResetToken: values.Mother().IDValid().Value(),
		ID:         values.Mother().IDValid().Value(),
		Email:      values.Mother().EmailValid().Value(),
		Username:   values.Mother().UsernameValid().Value(),
	}
}

var Mother = mother.New[m]
