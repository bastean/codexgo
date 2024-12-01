package login

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/hashes"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
)

type Case struct {
	repository.Repository
	hashes.Hasher
}

func (c *Case) Run(email *user.Email, username *user.Username, plain *user.PlainPassword) (*user.User, error) {
	criteria := new(repository.Criteria)

	switch {
	case email != nil:
		criteria.Email = email
	case username != nil:
		criteria.Username = username
	}

	aggregate, err := c.Repository.Search(criteria)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	err = hashes.IsPasswordInvalid(c.Hasher, aggregate.CipherPassword.Value, plain.Value)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return aggregate, nil
}
