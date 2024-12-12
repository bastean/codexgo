package create

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/hashes"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
)

type Case struct {
	hashes.Hasher
	repository.Repository
}

func (c *Case) Run(aggregate *user.User) error {
	hashed, err := c.Hasher.Hash(aggregate.PlainPassword.Value)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	aggregate.CipherPassword, err = user.NewCipherPassword(hashed)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	err = c.Repository.Create(aggregate)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}
