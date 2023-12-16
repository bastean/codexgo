package update

import (
	"github.com/bastean/codexgo/context/pkg/user/domain/aggregate"
	"github.com/bastean/codexgo/context/pkg/user/domain/repository"
)

type Update struct {
	Repository repository.Repository
}

func (update *Update) Run(user *aggregate.User) {
	update.Repository.Update(user)
}

func NewUpdate(repository repository.Repository) *Update {
	return &Update{
		repository,
	}
}
