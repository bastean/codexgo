package forgot

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mother"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

type m struct {
	*mother.Mother
}

func (m *m) CommandAttributesValid() *CommandAttributes {
	return &CommandAttributes{
		ResetToken: values.Mother().IDValid().Value(),
		Email:      values.Mother().EmailValid().Value(),
	}
}

var Mother = mother.New[m]
