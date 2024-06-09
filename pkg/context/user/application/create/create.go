package create

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/types"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
)

type Create struct {
	model.Repository
}

func (create *Create) Run(user *aggregate.User) (types.Empty, error) {
	err := create.Repository.Save(user)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return nil, nil
}
