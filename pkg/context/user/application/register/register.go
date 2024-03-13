package register

import (
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
)

type Register struct {
	model.Repository
}

func (register *Register) Run(user *aggregate.User) {
	register.Repository.Save(user)
}

func NewRegister(repository model.Repository) *Register {
	return &Register{
		repository,
	}
}
