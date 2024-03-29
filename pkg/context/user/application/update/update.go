package update

import (
	sharedValueObject "github.com/bastean/codexgo/pkg/context/shared/domain/valueObject"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/domain/service"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
)

type Update struct {
	model.Repository
	model.Hashing
}

func (update *Update) Run(userUpdate *Command) {
	idVO := sharedValueObject.NewId(userUpdate.Id)

	userRegistered := update.Repository.Search(model.RepositorySearchFilter{Id: idVO})

	service.IsPasswordInvalid(update.Hashing, userRegistered.Password.Value, userUpdate.Password)

	if userUpdate.Email != "" {
		userRegistered.Email = sharedValueObject.NewEmail(userUpdate.Email)
	}

	if userUpdate.Username != "" {
		userRegistered.Username = valueObject.NewUsername(userUpdate.Username)
	}

	if userUpdate.UpdatedPassword != "" {
		userRegistered.Password = valueObject.NewPassword(userUpdate.UpdatedPassword)
	}

	update.Repository.Update(userRegistered)
}

func NewUpdate(repository model.Repository, hashing model.Hashing) *Update {
	return &Update{
		repository,
		hashing,
	}
}
