package read

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/role"
)

type Case struct {
	role.Repository
}

func (c *Case) Run(attributes *QueryAttributes) (*user.User, error) {
	id, err := values.New[*user.ID](attributes.ID)

	if err != nil {
		return nil, errors.BubbleUp(err)
	}

	aggregate, err := c.Repository.Search(&user.Criteria{
		ID: id,
	})

	if err != nil {
		return nil, errors.BubbleUp(err)
	}

	return aggregate, nil
}
