package register

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errs"
	"github.com/bastean/codexgo/pkg/context/shared/domain/types"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
)

type Register struct {
	model.Repository
}

func (register *Register) Run(user *aggregate.User) (*types.Empty, error) {
	err := register.Repository.Save(user)

	if err != nil {
		return nil, errs.BubbleUp("Run", err)
	}

	return nil, nil
}
