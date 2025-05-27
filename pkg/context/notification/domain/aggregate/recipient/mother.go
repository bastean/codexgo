package recipient

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mother"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

type m struct {
	*mother.Mother
}

func (m *m) RecipientValid() *Recipient {
	user, err := New(&Required{
		ID:       values.Mother().IDValid().Value(),
		Email:    values.Mother().EmailValid().Value(),
		Username: values.Mother().UsernameValid().Value(),
	})

	if err != nil {
		errors.Panic(err)
	}

	return user
}

var Mother = mother.New[m]
