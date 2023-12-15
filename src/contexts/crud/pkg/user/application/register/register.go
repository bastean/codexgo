package register

import (
	"github.com/bastean/codexgo/context/pkg/user/domain/aggregate"
	"github.com/bastean/codexgo/context/pkg/user/domain/repository"
)

type Register struct {
	Repository repository.Repository
}

func (register *Register) Run(user *aggregate.User) {
	register.Repository.Save(user)
}

func NewRegister(repository repository.Repository) *Register {
	return &Register{
		repository,
	}
}
