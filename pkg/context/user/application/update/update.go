package update

import (
	sharedVO "github.com/bastean/codexgo/pkg/context/shared/domain/valueObject"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/domain/repository"
	"github.com/bastean/codexgo/pkg/context/user/domain/service"
	userVO "github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
)

type Update struct {
	repository.Repository
	model.Hashing
}

func (update *Update) Run(userUpdate Command) {
	idVO := sharedVO.NewId(userUpdate.Id)

	userRegistered := update.Repository.Search(repository.Filter{Id: idVO})

	service.IsPasswordInvalid(update.Hashing, userRegistered.Password.Value, userUpdate.Password)

	user := &aggregate.User{}

	user.Id = idVO

	if userUpdate.Email != "" {
		user.Email = sharedVO.NewEmail(userUpdate.Email)
	}

	if userUpdate.Username != "" {
		user.Username = userVO.NewUsername(userUpdate.Username)
	}

	if userUpdate.UpdatedPassword != "" {
		user.Password = userVO.NewPassword(userUpdate.UpdatedPassword)
	}

	update.Repository.Update(user)
}

func NewUpdate(repository repository.Repository, hashing model.Hashing) *Update {
	return &Update{
		repository,
		hashing,
	}
}
