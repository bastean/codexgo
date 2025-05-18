package verify

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mother"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
)

type m struct {
	*mother.Mother
}

func (m *m) CommandAttributesValid() *CommandAttributes {
	return &CommandAttributes{
		VerifyToken: user.Mother.IDValid().Value(),
		ID:          user.Mother.IDValid().Value(),
	}
}

var Mother = mother.New[m]()
