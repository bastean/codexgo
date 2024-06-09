package update

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/types"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/domain/service"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueobj"
)

type Update struct {
	model.Repository
	model.Hashing
}

func (update *Update) Run(userUpdate *Command) (types.Empty, error) {
	idVO, err := valueobj.NewId(userUpdate.Id)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	userRegistered, err := update.Repository.Search(&model.RepositorySearchCriteria{
		Id: idVO,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	err = service.IsPasswordInvalid(update.Hashing, userRegistered.Password.Value(), userUpdate.Password)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	var errEmail, errUsername, errPassword error

	if userUpdate.Email != "" {
		userRegistered.Email, errEmail = valueobj.NewEmail(userUpdate.Email)
	}

	if userUpdate.Username != "" {
		userRegistered.Username, errUsername = valueobj.NewUsername(userUpdate.Username)
	}

	if userUpdate.UpdatedPassword != "" {
		userRegistered.Password, errPassword = valueobj.NewPassword(userUpdate.UpdatedPassword)
	} else {
		userRegistered.Password = nil
	}

	err = errors.Join(errEmail, errUsername, errPassword)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	err = update.Repository.Update(userRegistered)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return nil, nil
}
