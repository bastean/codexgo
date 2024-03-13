package update

import (
	sharedVO "github.com/bastean/codexgo/pkg/context/shared/domain/valueObject"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/domain/service"
	userVO "github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
)

type Update struct {
	model.Repository
	model.Hashing
}

func (update *Update) Run(userUpdate *Command) {
	idVO := sharedVO.NewId(userUpdate.Id)

	userRegistered := update.Repository.Search(model.RepositorySearchFilter{Id: idVO})

	service.IsPasswordInvalid(update.Hashing, userRegistered.Password.Value, userUpdate.Password)

	if userUpdate.Email != "" {
		userRegistered.Email = sharedVO.NewEmail(userUpdate.Email)
	}

	if userUpdate.Username != "" {
		userRegistered.Username = userVO.NewUsername(userUpdate.Username)
	}

	if userUpdate.UpdatedPassword != "" {
		userRegistered.Password = userVO.NewPassword(userUpdate.UpdatedPassword)
	}

	update.Repository.Update(userRegistered)
}

func NewUpdate(repository model.Repository, hashing model.Hashing) *Update {
	return &Update{
		repository,
		hashing,
	}
}
