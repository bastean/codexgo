package create

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/serror"
	"github.com/bastean/codexgo/pkg/context/shared/domain/stype"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
)

type Create struct {
	model.Repository
}

func (create *Create) Run(user *aggregate.User) (*stype.Empty, error) {
	err := create.Repository.Save(user)

	if err != nil {
		return nil, serror.BubbleUp(err, "Run")
	}

	return nil, nil
}
