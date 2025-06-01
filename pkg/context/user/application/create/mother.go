package create

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mother"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
)

type m struct {
	*mother.Mother
}

func (m *m) CommandAttributesValid() *CommandAttributes {
	return &CommandAttributes{
		VerifyToken: values.Mother().IDValid().Value(),
		ID:          values.Mother().IDValid().Value(),
		Email:       values.Mother().EmailValid().Value(),
		Username:    values.Mother().UsernameValid().Value(),
		Password:    user.Mother().PlainPasswordValid().Value(),
	}
}

var Mother = mother.New[m]
