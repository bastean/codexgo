package login

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mother"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
)

type m struct {
	*mother.Mother
}

func (m *m) QueryAttributesValid() *QueryAttributes {
	return &QueryAttributes{
		Email:    user.Mother.EmailValid().Value(),
		Username: user.Mother.UsernameValid().Value(),
		Password: user.Mother.PlainPasswordValid().Value(),
	}
}

var Mother = mother.New[m]()
