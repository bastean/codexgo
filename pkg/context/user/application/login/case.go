package login

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/role"
)

type Case struct {
	role.Repository
	roles.Hasher
}

func (c *Case) Run(email *user.Email, username *user.Username, plain *user.PlainPassword) (*user.User, error) {
	aggregate, err := c.Repository.Search(&user.Criteria{
		Email:    email,
		Username: username,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	err = c.Hasher.Compare(aggregate.CipherPassword.Value, plain.Value)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return aggregate, nil
}
