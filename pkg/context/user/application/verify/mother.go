package verify

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mother"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

type m struct {
	*mother.Mother
}

func (m *m) CommandAttributesValid() *CommandAttributes {
	return &CommandAttributes{
		VerifyToken: values.Mother().TokenValid().Value(),
		ID:          values.Mother().IDValid().Value(),
	}
}

var Mother = mother.New[m]
