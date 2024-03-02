package register

import (
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/domain/repository"
)

type Register struct {
	repository.Repository
}

func (register *Register) Run(user *aggregate.User) {
	register.Repository.Save(user)
}

func NewRegister(repository repository.Repository) *Register {
	return &Register{
		repository,
	}
}
