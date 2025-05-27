package read

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mother"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

type m struct {
	*mother.Mother
}

func (m *m) QueryAttributesValid() *QueryAttributes {
	return &QueryAttributes{
		ID: values.Mother().IDValid().Value(),
	}
}

var Mother = mother.New[m]
