package login

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/hashes"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
)

type Case struct {
	repository.Repository
	hashes.Hashing
}

func (c *Case) Run(email *user.Email, plain *user.PlainPassword) (*user.User, error) {
	aggregate, err := c.Repository.Search(&repository.SearchCriteria{
		Email: email,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	err = hashes.IsPasswordInvalid(c.Hashing, aggregate.CipherPassword.Value, plain.Value)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return aggregate, nil
}
