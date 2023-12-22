package update

import (
	sharedVO "github.com/bastean/codexgo/context/pkg/shared/domain/valueObjects"
	"github.com/bastean/codexgo/context/pkg/user/domain/aggregate"
	"github.com/bastean/codexgo/context/pkg/user/domain/models"
	"github.com/bastean/codexgo/context/pkg/user/domain/repository"
	"github.com/bastean/codexgo/context/pkg/user/domain/services"
	userVO "github.com/bastean/codexgo/context/pkg/user/domain/valueObjects"
)

type Update struct {
	Repository repository.Repository
	Hashing    models.Hashing
}

func (update *Update) Run(userUpdate Command) {
	idVO := sharedVO.NewId(userUpdate.Id)

	userRegistered := update.Repository.Search(repository.Filter{Id: idVO})

	services.IsPasswordInvalid(update.Hashing, userRegistered.Password.Value, userUpdate.Password)

	user := &aggregate.User{}

	user.Id = sharedVO.NewId(userUpdate.Id)

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

func NewUpdate(repository repository.Repository, hashing models.Hashing) *Update {
	return &Update{
		repository,
		hashing,
	}
}
