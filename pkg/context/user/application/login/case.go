package login

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/role"
)

type Case struct {
	role.Repository
	roles.Hasher
}

func (c *Case) Run(attributes *QueryAttributes) (*user.User, error) {
	if attributes.Email == "" && attributes.Username == "" {
		return nil, errors.New[errors.Failure](&errors.Bubble{
			What: "Email or Username required",
		})
	}

	var (
		err      error
		email    *user.Email
		username *user.Username
	)

	if attributes.Email != "" {
		email, err = values.New[*user.Email](attributes.Email)
	}

	if err != nil {
		return nil, errors.BubbleUp(err)
	}

	if attributes.Username != "" {
		username, err = values.New[*user.Username](attributes.Username)
	}

	if err != nil {
		return nil, errors.BubbleUp(err)
	}

	plainPassword, err := values.New[*user.PlainPassword](attributes.Password)

	if err != nil {
		return nil, errors.BubbleUp(err)
	}

	aggregate, err := c.Repository.Search(&user.Criteria{
		Email:    email,
		Username: username,
	})

	if err != nil {
		return nil, errors.BubbleUp(err)
	}

	err = c.Hasher.Compare(aggregate.Password.Value(), plainPassword.Value())

	if err != nil {
		return nil, errors.BubbleUp(err)
	}

	return aggregate, nil
}
