package verify

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/role"
)

type Case struct {
	role.Repository
}

func (c *Case) Run(verify, id *user.ID) error {
	aggregate, err := c.Repository.Search(&user.Criteria{
		ID: id,
	})

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	if aggregate.IsVerified() {
		return nil
	}

	err = aggregate.ValidateVerify(verify)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	aggregate.Verified, err = user.NewVerified(true)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	aggregate.Verify = nil

	err = c.Repository.Update(aggregate)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}
