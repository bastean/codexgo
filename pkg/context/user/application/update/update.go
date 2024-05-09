package update

import (
	"errors"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errs"
	"github.com/bastean/codexgo/pkg/context/shared/domain/types"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/domain/service"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
)

type Update struct {
	model.Repository
	model.Hashing
}

func (update *Update) Run(userUpdate *Command) (*types.Empty, error) {
	idVO, err := valueObject.NewId(userUpdate.Id)

	if err != nil {
		return nil, errs.BubbleUp(err, "Run")
	}

	userRegistered, err := update.Repository.Search(model.RepositorySearchCriteria{
		Id: idVO,
	})

	if err != nil {
		return nil, errs.BubbleUp(err, "Run")
	}

	err = service.IsPasswordInvalid(update.Hashing, userRegistered.Password.Value(), userUpdate.Password)

	if err != nil {
		return nil, errs.BubbleUp(err, "Run")
	}

	var emailErr, usernameErr, passwordErr error

	if userUpdate.Email != "" {
		userRegistered.Email, emailErr = valueObject.NewEmail(userUpdate.Email)
	}

	if userUpdate.Username != "" {
		userRegistered.Username, usernameErr = valueObject.NewUsername(userUpdate.Username)
	}

	if userUpdate.UpdatedPassword != "" {
		userRegistered.Password, passwordErr = valueObject.NewPassword(userUpdate.UpdatedPassword)
	}

	err = errors.Join(emailErr, usernameErr, passwordErr)

	if err != nil {
		return nil, errs.BubbleUp(err, "Run")
	}

	err = update.Repository.Update(userRegistered)

	if err != nil {
		return nil, errs.BubbleUp(err, "Run")
	}

	return nil, nil
}
